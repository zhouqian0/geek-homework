package service

import (
	"context"
	v1 "dipper/api/server/service/v1"
	"dipper/app/server/service/internal/biz"
	"github.com/go-kratos/kratos/v2/errors"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/peer"
)

func (s *ServerService) CreateHost(ctx context.Context, req *v1.CreateHostReq) (*v1.CreateHostReply, error) {
	// todo 参数校验器 https://go-kratos.dev/docs/component/middleware/validate
	rv, duplicated, err := s.host.Create(ctx, &biz.Host{
		Name:       req.GetName(),
		Manager:    req.GetManager(),
		Phone:      req.GetPhone(),
		VerifyCode: req.GetVerifyCode(),
	})
	if err != nil {
		s.log.Errorf("[CreateHost] failed to create host: %v", err)
		return &v1.CreateHostReply{
			Code: codeFailed,
			Msg:  msgFailed,
		}, nil
	}

	if duplicated {
		return &v1.CreateHostReply{
			Code: codeConflict,
			Msg:  msgDuplicated,
		}, nil
	}
	return &v1.CreateHostReply{
		Code: codeSuccess,
		Msg:  msgSuccess,
		Data: &v1.HostInfo{
			Id:         rv.ID,
			Name:       rv.Name,
			Manager:    rv.Manager,
			Phone:      rv.Phone,
			VerifyCode: rv.VerifyCode,
		},
	}, nil
}

func (s *ServerService) ListHost(ctx context.Context, _ *v1.ListHostReq) (*v1.ListHostReply, error) {
	rv, err := s.host.List(ctx)
	if err != nil {
		s.log.Errorf("[ListHost] failed to list host: %v", err)
		return &v1.ListHostReply{
			Code: codeFailed,
			Msg:  msgFailed,
		}, nil
	}

	rs := make([]*v1.HostInfo, len(rv))
	for i, x := range rv {
		rs[i] = &v1.HostInfo{
			Id:      x.ID,
			Name:    x.Name,
			Manager: x.Manager,
			Phone:   x.Phone,
		}
	}
	return &v1.ListHostReply{
		Code: codeSuccess,
		Msg:  msgSuccess,
		Data: rs,
	}, nil
}

func (s *ServerService) GetHostByVerifyCode(ctx context.Context, req *v1.GetHostByVerifyCodeReq) (*v1.GetHostByVerifyCodeReply, error) {
	if !s.checkClientCertValid(ctx) {
		return nil, errors.Unauthorized("", "")
	}

	rv, has, err := s.host.GetHostByVerifyCode(ctx, req.GetVerifyCode())
	if err != nil {
		s.log.Errorf("[GetHostByVerifyCode] failed to get host by verify code: %s", err)
		return nil, err
	}

	rs := &v1.HostInfo{}
	if has {
		rs.Name = rv.Name
		rs.Manager = rv.Manager
		rs.Phone = rv.Phone
	}
	return &v1.GetHostByVerifyCodeReply{
		Host: rs,
	}, nil
}

func (s *ServerService) AuthHost(ctx context.Context, req *v1.AuthHostReq) (*v1.AuthHostReply, error) {
	blocked, err := s.blacklist.CheckCodeBlacklisted(ctx, req.GetVerifyCode())
	if err != nil {
		s.log.Errorf("[AuthHost] failed to check verify code blacklisted: %v", err)
		return nil, err
	}
	if blocked {
		return nil, errors.Unauthorized("", "")
	}

	rv, has, err := s.host.GetHostByVerifyCode(ctx, req.GetVerifyCode())
	if err != nil {
		s.log.Errorf("[AuthHost] failed to get host by verify code: %v", err)
		return nil, err
	}
	if !has {
		return nil, errors.Unauthorized("", "")
	}

	cert, err := s.host.GenHostCert(ctx)
	if err != nil {
		s.log.Errorf("[AuthHost] failed to gen cert, %v", err)
		return nil, err
	}
	_, err = s.host.UpdateCertNumByHostID(ctx, rv.ID, cert.Num)
	if err != nil {
		s.log.Errorf("[AuthHost] failed to save cert num, %v", err)
		return nil, err
	}
	return &v1.AuthHostReply{
		Cert: string(cert.Cert),
		Key:  string(cert.Key),
	}, nil
}

func (s *ServerService) checkClientCertValid(ctx context.Context) bool {
	p, ok := peer.FromContext(ctx)
	if !ok {
		return false
	}
	tlsInfo := p.AuthInfo.(credentials.TLSInfo)
	if len(tlsInfo.State.PeerCertificates) == 0 {
		return false
	}
	return true
}
