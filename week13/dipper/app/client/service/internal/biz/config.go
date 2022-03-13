package biz

import (
	"context"
	"dipper/app/client/service/internal/biz/dto"
	"encoding/json"
	"github.com/go-kratos/kratos/v2/log"
)

const (
	KeyCertPem = "cert.pem"
	keyCertKey = "cert.key"
	keyHost    = "host.info"
)

type Config struct {
	Key   string
	Value string
}

type ConfigRepo interface {
	CreatConfig(ctx context.Context, c *Config) (*Config, error)
	BatchCreatConfig(ctx context.Context, cfgs ...*Config) ([]*Config, error)
	DeleteConfig(ctx context.Context, key string) error
	BatchDeleteConfig(ctx context.Context, keys ...string) error
	UpdateConfig(ctx context.Context, c *Config) (*Config, error)
	BatchUpdateConfig(ctx context.Context, cfgs ...*Config) ([]*Config, error)
	GetConfig(ctx context.Context, key string) (*Config, bool, error)
	BatchGetConfig(ctx context.Context, keys ...string) ([]*Config, error)

	UpdateCache(key string, value []byte) error
	DeleteCache(key string) error
	GetCache(key string) ([]byte, bool, error)

	IsDBNotFound(err error) bool
	IsCacheNotFound(err error) bool
}

type ConfigUseCase struct {
	repo ConfigRepo
	log  *log.Helper
}

func NewConfigUseCase(repo ConfigRepo, logger log.Logger) *ConfigUseCase {
	return &ConfigUseCase{repo: repo, log: log.NewHelper(log.With(logger, "module", "usecase/config"))}
}

func (uc *ConfigUseCase) GetCert(ctx context.Context) (*dto.Cert, bool, error) {
	var (
		certPem, certKey []byte
		err              error
	)
	if certPem, _, err = uc.repo.GetCache(KeyCertPem); err != nil {
		return nil, false, err
	}
	if certKey, _, err = uc.repo.GetCache(keyCertKey); err != nil {
		return nil, false, err
	}
	if len(certPem) != 0 && len(certKey) != 0 {
		return &dto.Cert{
			Cert: certPem,
			Key:  certKey,
		}, true, nil
	} else if err = uc.removeCertFromCache(); err != nil {
		return nil, false, err
	}

	rvs, err := uc.repo.BatchGetConfig(ctx, KeyCertPem, keyCertKey)
	if err != nil {
		return nil, false, err
	}
	for _, v := range rvs {
		if v.Key == KeyCertPem {
			certPem = []byte(v.Value)
		} else if v.Key == keyCertKey {
			certKey = []byte(v.Value)
		}
	}
	if len(certPem) > 0 && len(certKey) > 0 {
		_ = uc.saveCertToCache(certPem, certKey)
		return &dto.Cert{
			Cert: certPem,
			Key:  certKey,
		}, true, nil
	}
	return &dto.Cert{}, false, nil
}

func (uc *ConfigUseCase) UpdateCert(ctx context.Context, certPem, certKey []byte) error {
	var err error
	if err = uc.removeCertFromCache(); err != nil {
		return err
	}
	if err = uc.repo.BatchDeleteConfig(ctx, KeyCertPem, keyCertKey); err != nil && !uc.repo.IsDBNotFound(err) {
		return err
	}

	if _, err = uc.repo.BatchCreatConfig(ctx,
		&Config{
			Key:   KeyCertPem,
			Value: string(certPem)},
		&Config{
			Key:   keyCertKey,
			Value: string(certKey)}); err != nil {
		return err
	}
	_ = uc.saveCertToCache(certPem, certKey)
	return nil
}

func (uc *ConfigUseCase) saveCertToCache(certPem, certKey []byte) error {
	var err error
	if err = uc.repo.UpdateCache(KeyCertPem, certPem); err != nil {
		return err
	}
	if err = uc.repo.UpdateCache(keyCertKey, certKey); err != nil {
		return err
	}
	return nil
}

func (uc *ConfigUseCase) removeCertFromCache() error {
	var err error
	if err = uc.repo.DeleteCache(KeyCertPem); err != nil && !uc.repo.IsCacheNotFound(err) {
		return err
	}
	if err = uc.repo.DeleteCache(keyCertKey); err != nil && !uc.repo.IsCacheNotFound(err) {
		return err
	}
	return nil
}

func (uc *ConfigUseCase) UpdateHost(_ context.Context, host *dto.Host) error {
	b, _ := json.Marshal(host)
	return uc.repo.UpdateCache(keyHost, b)
}

func (uc *ConfigUseCase) GetHost(_ context.Context) (*dto.Host, bool, error) {
	rv, has, err := uc.repo.GetCache(keyHost)
	if err != nil {
		return nil, false, err
	}
	if !has {
		return nil, false, nil
	}

	host := &dto.Host{}
	if err = json.Unmarshal(rv, host); err != nil {
		return nil, false, err
	}
	return host, true, nil
}
