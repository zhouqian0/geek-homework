package machine

import (
	"bytes"
	"net"
	"os"
	"os/exec"
	"syscall"
	"unsafe"
)

type Net struct {
	Name       string `json:"name"`        // 名称
	IP         string `json:"ip"`          // ip地址
	SubnetMask string `json:"subnet_mask"` // 子网掩码
	Gateway    string `json:"gateway"`     // 默认网关
	MAC        string `json:"mac"`         // 物理地址
}

// ipAdapter 网络适配器信息。
type ipAdapter struct {
	ip         string // ip 地址
	subnetMask string // 子网掩码
	gateway    string // 网关地址
}

func GetNet() ([]*Net, error) {
	// 获取不同网卡的默认网关信息集合
	adapterMp, err := pkgIPAdapterMap()
	if err != nil {
		return nil, err
	}

	// 获取网络接口
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	nets := make([]*Net, 0, len(ifaces))
	// 遍历查找每个网卡信息
	for _, iface := range ifaces {
		// 忽略虚拟网卡（没有 mac 地址的网卡）
		if iface.HardwareAddr == nil {
			continue
		}

		// 查找该网卡对应的网络适配器信息
		adapter := adapterMp[iface.Index]
		if adapter == nil {
			continue
		}
		// 封装网卡对象
		nets = append(nets, &Net{
			Name:       iface.Name,
			IP:         adapter.ip,
			SubnetMask: adapter.subnetMask,
			Gateway:    adapter.gateway,
			MAC:        iface.HardwareAddr.String(),
		})
	}
	return nets, nil
}

func UpdateNet(name, ip, subnetMask, gateway string) error {
	cmd := exec.Command("netsh",
		"interface", "ip", "set", "address", name, "static", ip, subnetMask, gateway)
	if _, err := cmd.Output(); err != nil {
		return err
	}
	return nil
}

// todo 无法区分以太网适配器之后的诸如蓝牙，虚拟网卡，物理网卡的细类。ps: 硬要做也只能匹配字符串了。
// pkgIPAdapterMap 获取网络适配器网卡信息，map key 为网卡在系统中的下标，val 为对应 ip，掩码和网关。
func pkgIPAdapterMap() (map[int]*ipAdapter, error) {
	// 获取网络适配器信息。
	a, err := getIPAdapterInfo()
	if err != nil {
		return nil, err
	}

	// 无线网卡类型
	typeWireless := uint32(71)
	ans := make(map[int]*ipAdapter)
	for ; a != nil; a = a.Next {
		// 排除无线网卡
		if a.Type == typeWireless {
			continue
		}

		adapter := &ipAdapter{}

		// 获取网关信息
		gateway := &a.GatewayList
		for ; gateway != nil; gateway = gateway.Next {
			adapter.gateway = string(bytes.Trim(gateway.IpAddress.String[:], "\x00"))
		}

		// 获取网卡和子网掩码信息
		ipAddress := &a.IpAddressList
		for ; ipAddress != nil; ipAddress = ipAddress.Next {
			adapter.ip = string(bytes.Trim(ipAddress.IpAddress.String[:], "\x00"))
			adapter.subnetMask = string(bytes.Trim(ipAddress.IpMask.String[:], "\x00"))
		}
		ans[int(a.Index)] = adapter
	}
	return ans, nil
}

// getIPAdapterInfo 获取系统网络适配器信息。
func getIPAdapterInfo() (*syscall.IpAdapterInfo, error) {
	b := make([]byte, 1000)
	l := uint32(len(b))
	a := (*syscall.IpAdapterInfo)(unsafe.Pointer(&b[0]))
	// 通过 syscall 调用系统底层获取网络适配器信息
	err := syscall.GetAdaptersInfo(a, &l)
	if err == syscall.ERROR_BUFFER_OVERFLOW {
		b = make([]byte, l)
		a = (*syscall.IpAdapterInfo)(unsafe.Pointer(&b[0]))
		err = syscall.GetAdaptersInfo(a, &l)
	}
	if err != nil {
		return nil, os.NewSyscallError("GetAdaptersInfo", err)
	}
	return a, nil
}
