package machine

import (
	"fmt"
	"golang.org/x/sys/windows/registry"
	"syscall"
)

func GetComputerName() (string, error) {
	var (
		size   uint32 = 128
		buffer        = make([]uint16, size)
	)
	domain, err := syscall.UTF16PtrFromString("COMPUTERNAME")
	if err != nil {
		return "", err
	}
	n, err := syscall.GetEnvironmentVariable(domain, &buffer[0], size)
	if err != nil {
		return "", err
	}
	return syscall.UTF16ToString(buffer[:n]), nil
}

func GetWindowsProductName() (string, error) {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE,
		`SOFTWARE\Microsoft\Windows NT\CurrentVersion`, registry.QUERY_VALUE)
	if err != nil {
		return "", err
	}
	defer func(k registry.Key) {
		_ = k.Close()
	}(k)

	pn, _, err := k.GetStringValue("ProductName")
	if err != nil {
		return "", err
	}
	bit := 32 << (^uint(0) >> 63)
	return fmt.Sprintf("%s(%d)", pn, bit), nil
}

func GetWindowsDisplayVersion() (string, error) {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE,
		`SOFTWARE\Microsoft\Windows NT\CurrentVersion`, registry.QUERY_VALUE)
	if err != nil {
		return "", err
	}
	defer func(k registry.Key) {
		_ = k.Close()
	}(k)

	dv, _, err := k.GetStringValue("DisplayVersion")
	if err != nil {
		return "", err
	}
	return dv, nil
}
