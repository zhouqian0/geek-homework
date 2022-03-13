package util

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"io/ioutil"
	"math/big"
	"net"
	"time"
)

func GenCertPool(filePath string) (*x509.CertPool, error) {
	ca, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	pool := x509.NewCertPool()
	pool.AppendCertsFromPEM(ca)
	return pool, nil
}

// GenClient 通过父证书和密钥生成客户端证书和密钥。
func GenClient(pemPath, keyPath string, uid int64) ([]byte, []byte, error) {
	parentPEM, err := parsePemFile(pemPath)
	if err != nil {
		return nil, nil, err
	}
	parentKey, err := parseKeyFile(keyPath)
	if err != nil {
		return nil, nil, err
	}

	return genChild(parentPEM, parentKey, uid, "")
}

// genChild 生成子证书。
func genChild(parentPEM *x509.Certificate, parentKey *rsa.PrivateKey, uid int64, serverAddr string) ([]byte, []byte, error) {
	cert := &x509.Certificate{
		SerialNumber: big.NewInt(uid), // 唯一 id，用作标识证书，建议使用统一发号器，避免生成重复 id
		Subject:      pkix.Name{},
		NotBefore:    time.Now(),                   // 证书即刻生效
		NotAfter:     time.Now().AddDate(50, 0, 0), // 证书有效期 50 年
		SubjectKeyId: []byte{1, 2, 3, 4, 6},
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:     x509.KeyUsageDigitalSignature,
	}
	// 如果服务地址不为空则为证书设置 ip 地址
	if len(serverAddr) > 0 {
		cert.IPAddresses = []net.IP{net.ParseIP(serverAddr)}
	}

	// 生成证书密钥对
	key, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return nil, nil, err
	}
	// 生成证书
	certBytes, err := x509.CreateCertificate(rand.Reader, cert, parentPEM, &key.PublicKey, parentKey)
	if err != nil {
		return nil, nil, err
	}

	// 打包证书
	certPEM := new(bytes.Buffer)
	_ = pem.Encode(certPEM, &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: certBytes,
	})
	// 打包密钥
	certKey := new(bytes.Buffer)
	_ = pem.Encode(certKey, &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	})

	return certPEM.Bytes(), certKey.Bytes(), nil
}

func parsePemFile(path string) (*x509.Certificate, error) {
	certPEMBlock, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	//获取证书信息 -----BEGIN CERTIFICATE-----   -----END CERTIFICATE-----
	//这里返回的第二个值是证书中剩余的 block, 一般是rsa私钥 也就是 -----BEGIN RSA PRIVATE KEY 部分
	//一般证书的有效期，组织信息等都在第一个部分里
	certDERBlock, _ := pem.Decode(certPEMBlock)
	if certDERBlock == nil {
		return nil, nil
	}

	cert, err := x509.ParseCertificate(certDERBlock.Bytes)
	if err != nil {
		return nil, err
	}
	return cert, nil
}

func parseKeyFile(path string) (*rsa.PrivateKey, error) {
	certPEMBlock, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	//获取证书信息 -----BEGIN CERTIFICATE-----   -----END CERTIFICATE-----
	//这里返回的第二个值是证书中剩余的 block, 一般是rsa私钥 也就是 -----BEGIN RSA PRIVATE KEY 部分
	//一般证书的有效期，组织信息等都在第一个部分里
	certDERBlock, _ := pem.Decode(certPEMBlock)
	if certDERBlock == nil {
		return nil, nil
	}

	key, err := x509.ParsePKCS1PrivateKey(certDERBlock.Bytes)
	if err != nil {
		return nil, err
	}
	return key, nil
}
