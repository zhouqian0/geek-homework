package machine

import (
	"crypto/md5"
	"fmt"
	"net"
	"os/exec"
	"regexp"
)

func GetMachineUniqueID() (string, error) {
	mac, err := getMac()
	if err != nil {
		return "", err
	}
	cpuID, err := getCpuID()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%X", md5.Sum([]byte(mac+cpuID))), nil
}

func getMac() (string, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}

	inter := interfaces[0]
	mac := inter.HardwareAddr.String()
	return mac, nil
}

func getCpuID() (string, error) {
	cmd := exec.Command("wmic", "cpu", "get", "ProcessorID")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	str, reg := string(out), regexp.MustCompile("\\s+")
	str = reg.ReplaceAllString(str, "")
	startIdx := 11
	return str[startIdx:], nil
}
