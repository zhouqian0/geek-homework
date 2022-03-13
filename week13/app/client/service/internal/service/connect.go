package service

import (
	"context"
	v1 "dipper/api/client/service/v1"
)

func (s *ClientService) CheckFirstConn(ctx context.Context, _ *v1.CheckFirstConnReq) (*v1.CheckFirstConnReply, error) {
	_, has, err := s.config.GetCert(ctx)
	if err != nil {
		s.log.Errorf("[CheckConn] failed to get cert: %v", err)
		return &v1.CheckFirstConnReply{
			Code: codeFailed,
			Msg:  msgFailed,
		}, nil
	}
	if has {
		return &v1.CheckFirstConnReply{
			Code: codeSuccess,
			Msg:  msgSuccess,
			Data: &v1.CheckFirstConnReply_ConnCheck{
				First: false,
			},
		}, nil
	}

	code, err := s.machine.GetMachineUniqueID()
	if err != nil {
		s.log.Errorf("[CheckConn] failed to gen machine unique id: %v", err)
		return &v1.CheckFirstConnReply{
			Code: codeFailed,
			Msg:  msgFailed,
		}, nil
	}
	return &v1.CheckFirstConnReply{
		Code: codeSuccess,
		Msg:  msgSuccess,
		Data: &v1.CheckFirstConnReply_ConnCheck{
			First:      true,
			VerifyCode: code,
		},
	}, nil
}

func (s *ClientService) ConnectServer(ctx context.Context, _ *v1.ConnectServerReq) (*v1.ConnectServerReply, error) {
	verifyCode, err := s.machine.GetMachineUniqueID()
	if err != nil {
		s.log.Errorf("[ConnectServer] failed to get machine unique id: %v", err)
		return &v1.ConnectServerReply{
			Code: codeFailed,
			Msg:  msgFailed,
		}, nil
	}

	cert, has, err := s.config.GetCert(ctx)
	if err != nil {
		s.log.Errorf("[ConnectServer] failed to get cert: %v", err)
		return &v1.ConnectServerReply{
			Code: codeFailed,
			Msg:  msgFailed,
		}, nil
	}
	if !has {
		authorized := false
		if cert, authorized, err = s.host.GetCert(ctx, verifyCode); err != nil {
			s.log.Errorf("[ConnectServer] failed to server auth: %v", err)
			return &v1.ConnectServerReply{
				Code: codeFailed,
				Msg:  msgFailed,
			}, nil
		}
		if !authorized {
			return &v1.ConnectServerReply{
				Code: codeUnauthorized,
				Msg:  msgUnauthorized,
			}, nil
		}
		if err = s.config.UpdateCert(ctx, cert.Cert, cert.Key); err != nil {
			s.log.Errorf("[ConnectServer] failed to save cert: %v", err)
			return &v1.ConnectServerReply{
				Code: codeFailed,
				Msg:  msgFailed,
			}, nil
		}
	}

	if err = s.host.UpdateHostServiceClient(ctx, cert.Cert, cert.Key); err != nil {
		s.log.Errorf("[ConnectServer] failed to update host client: %v", err)
		return &v1.ConnectServerReply{
			Code: codeFailed,
			Msg:  msgFailed,
		}, nil
	}

	host, err := s.host.GetHost(ctx, verifyCode)
	if err != nil {
		s.log.Errorf("[ConnectServer] failed to get host info: %v", err)
		return &v1.ConnectServerReply{
			Code: codeFailed,
			Msg:  msgFailed,
		}, nil
	}
	if err = s.config.UpdateHost(ctx, host); err != nil {
		s.log.Errorf("[ConnectServer] failed to save host info: %v", err)
		return &v1.ConnectServerReply{
			Code: codeFailed,
			Msg:  msgFailed,
		}, nil
	}

	return &v1.ConnectServerReply{
		Code: codeSuccess,
		Msg:  msgSuccess,
	}, nil
}
