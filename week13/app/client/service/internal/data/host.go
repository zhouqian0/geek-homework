package data

import (
	"context"
	serverv1 "dipper/api/server/service/v1"
	"dipper/app/client/service/internal/biz"
	"dipper/app/client/service/internal/biz/dto"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

type HostRepo struct {
	data *Data
	log  *log.Helper
}

func NewHostRepo(data *Data, logger log.Logger) biz.HostRepo {
	return &HostRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/host")),
	}
}

func (r *HostRepo) GetServerAuth(ctx context.Context, code string) (*dto.Cert, bool, error) {
	cert, err := r.data.hostServer.client.AuthHost(ctx, &serverv1.AuthHostReq{VerifyCode: code})
	if err != nil {
		if errors.IsUnauthorized(err) {
			return nil, false, nil
		}
		return nil, false, err
	}
	return &dto.Cert{
		Cert: []byte(cert.GetCert()),
		Key:  []byte(cert.GetKey()),
	}, true, nil
}

func (r *HostRepo) UpdateHostServiceClient(_ context.Context, cert, key []byte) error {
	return updateHostServer(r.data.hostServer, cert, key)
}

func (r *HostRepo) GetHostInfo(ctx context.Context, code string) (*dto.Host, error) {
	rv, err := r.data.hostServer.client.GetHostByVerifyCode(ctx, &serverv1.GetHostByVerifyCodeReq{VerifyCode: code})
	if err != nil {
		return nil, err
	}
	if rv.GetHost() == nil {
		return &dto.Host{}, nil
	}
	return &dto.Host{
		Name:    rv.GetHost().GetName(),
		Manager: rv.GetHost().GetManager(),
		Phone:   rv.GetHost().GetPhone(),
	}, nil
}
