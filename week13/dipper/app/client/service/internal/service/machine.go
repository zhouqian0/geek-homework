package service

import (
	"context"
	v1 "dipper/api/client/service/v1"
)

func (s *ClientService) ListNet(context.Context, *v1.ListNetReq) (*v1.ListNetReply, error) {
	rv, err := s.machine.ListNet()
	if err != nil {
		s.log.Errorf("[ListNet] failed to list net: %v", err)
		return &v1.ListNetReply{
			Code: codeFailed,
			Msg:  msgFailed,
		}, nil
	}

	rs := make([]*v1.ListNetReply_Net, len(rv))
	for i, x := range rv {
		rs[i] = &v1.ListNetReply_Net{
			Name:       x.Name,
			Ip:         x.IP,
			Gateway:    x.Gateway,
			SubnetMask: x.SubnetMask,
			Mac:        x.MAC,
		}
	}
	return &v1.ListNetReply{
		Code: codeSuccess,
		Msg:  msgSuccess,
		Data: rs,
	}, nil
}
func (s *ClientService) ListCom(context.Context, *v1.ListComReq) (*v1.ListComReply, error) {
	rv, err := s.machine.ListCOM()
	if err != nil {
		s.log.Errorf("[ListNet] failed to list com: %v", err)
		return &v1.ListComReply{
			Code: codeFailed,
			Msg:  msgFailed,
		}, nil
	}

	rs := make([]*v1.ListComReply_Com, len(rv))
	for i, x := range rv {
		rs[i] = &v1.ListComReply_Com{
			Name: x.Name,
			Desc: x.Desc,
		}
	}
	return &v1.ListComReply{
		Code: codeSuccess,
		Msg:  msgSuccess,
		Data: rs,
	}, nil
}
func (s *ClientService) GetSystemInfo(context.Context, *v1.GetSystemInfoReq) (*v1.GetSystemInfoReply, error) {
	x, err := s.machine.GetSystemInfo()
	if err != nil {
		s.log.Errorf("[GetSystemInfo] failed to get system info: %v", err)
		return &v1.GetSystemInfoReply{
			Code: codeFailed,
			Msg:  msgFailed,
		}, nil
	}

	return &v1.GetSystemInfoReply{
		Code: codeSuccess,
		Msg:  msgSuccess,
		Data: &v1.GetSystemInfoReply_SystemInfo{
			Name:    x.Name,
			Os:      x.Os,
			Version: x.Version,
		},
	}, nil
}
func (s *ClientService) ListAntivirus(context.Context, *v1.ListAntivirusReq) (*v1.ListAntivirusReply, error) {
	x, err := s.machine.ListAntivirus()
	if err != nil {
		s.log.Errorf("[ListAntivirus] failed to list anti-virus: %v", err)
		return &v1.ListAntivirusReply{
			Code: codeFailed,
			Msg:  msgFailed,
		}, nil
	}

	return &v1.ListAntivirusReply{
		Code: codeSuccess,
		Msg:  msgSuccess,
		Data: x,
	}, nil
}
