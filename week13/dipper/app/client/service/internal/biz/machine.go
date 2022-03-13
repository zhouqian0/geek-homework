package biz

import (
	"dipper/app/client/service/pkg/machine"
	"github.com/go-kratos/kratos/v2/log"
)

type Net struct {
	Name       string `json:"name"`       // 名称
	IP         string `json:"ip"`         // ip地址
	SubnetMask string `json:"subnetMask"` // 子网掩码
	Gateway    string `json:"gateway"`    // 默认网关
	MAC        string `json:"mac"`        // 物理地址
}

type COM struct {
	Name string `json:"name"` // 端口名称
	Desc string `json:"desc"` // 端口描述
}

type SystemInfo struct {
	Name    string `json:"name"`    // 电脑名称
	Os      string `json:"os"`      // 系统类型
	Version string `json:"version"` // 系统版本
}

type MachineUseCase struct {
	log *log.Helper
}

func NewMachineUseCase(logger log.Logger) *MachineUseCase {
	return &MachineUseCase{log: log.NewHelper(log.With(logger, "module", "usecase/machine"))}
}

func (uc *MachineUseCase) GetMachineUniqueID() (string, error) {
	uid, err := machine.GetMachineUniqueID()
	if err != nil {
		return "", err
	}

	stdLen := 12
	return uid[:stdLen], nil
}

func (uc *MachineUseCase) ListNet() ([]*Net, error) {
	rv, err := machine.GetNet()
	if err != nil {
		return nil, err
	}

	nets := make([]*Net, 0, len(rv))
	for _, x := range rv {
		nets = append(nets, &Net{
			Name:       x.Name,
			IP:         x.IP,
			SubnetMask: x.SubnetMask,
			Gateway:    x.Gateway,
			MAC:        x.MAC,
		})
	}
	return nets, nil
}

func (uc *MachineUseCase) UpdateNet(net *Net) error {
	return machine.UpdateNet(net.Name, net.IP, net.SubnetMask, net.Gateway)
}

func (uc *MachineUseCase) ListCOM() ([]*COM, error) {
	rv, err := machine.GetCOMs()
	if err != nil {
		return nil, err
	}

	coms := make([]*COM, 0, len(rv))
	for _, x := range rv {
		coms = append(coms, &COM{
			Name: x.Name,
			Desc: x.Description,
		})
	}
	return coms, nil
}

func (uc *MachineUseCase) GetSystemInfo() (*SystemInfo, error) {
	var (
		err  error
		info = &SystemInfo{}
	)
	if info.Name, err = machine.GetComputerName(); err != nil {
		return nil, err
	}
	if info.Os, err = machine.GetWindowsProductName(); err != nil {
		return nil, err
	}
	if info.Version, err = machine.GetWindowsDisplayVersion(); err != nil {
		return nil, err
	}
	return info, nil
}

func (uc *MachineUseCase) ListAntivirus() ([]string, error) {
	rv, err := machine.GetAntiVirusProducts()
	if err != nil {
		return nil, err
	}

	antivirus := make([]string, len(rv))
	for i, x := range rv {
		antivirus[i] = x.DisplayName
	}
	return antivirus, nil
}
