package service

import (
	"context"
	v1 "dipper/api/client/service/v1"
)

func (s *ClientService) GetHostInfo(ctx context.Context, _ *v1.GetHostInfoReq) (*v1.GetHostInfoReply, error) {
	rv, has, err := s.config.GetHost(ctx)
	if err != nil {
		s.log.Errorf("[GetHostInfo] failed to get host info: %v", err)
		return &v1.GetHostInfoReply{
			Code: codeFailed,
			Msg:  msgFailed,
		}, nil
	}

	rs := &v1.GetHostInfoReply_HostInfo{}
	if has {
		rs.Name = rv.Name
		rs.Manager = rv.Manager
		rs.Phone = rv.Phone
	}
	return &v1.GetHostInfoReply{
		Code: codeSuccess,
		Msg:  msgSuccess,
		Data: rs,
	}, nil
}
