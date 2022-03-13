package biz

import (
	"context"
	"dipper/app/client/service/internal/biz/dto"
	"github.com/go-kratos/kratos/v2/log"
)

type HostRepo interface {
	GetHostInfo(ctx context.Context, code string) (*dto.Host, error)
	GetServerAuth(ctx context.Context, code string) (*dto.Cert, bool, error)
	UpdateHostServiceClient(ctx context.Context, cert, key []byte) error
}

type HostUseCase struct {
	repo HostRepo
	log  *log.Helper
}

func NewHostUseCase(repo HostRepo, logger log.Logger) *HostUseCase {
	return &HostUseCase{repo: repo, log: log.NewHelper(log.With(logger, "module", "usecase/host"))}
}

func (uc *HostUseCase) GetCert(ctx context.Context, verifyCode string) (*dto.Cert, bool, error) {
	return uc.repo.GetServerAuth(ctx, verifyCode)
}

func (uc *HostUseCase) UpdateHostServiceClient(ctx context.Context, cert, key []byte) error {
	return uc.repo.UpdateHostServiceClient(ctx, cert, key)
}

func (uc *HostUseCase) GetHost(ctx context.Context, verifyCode string) (*dto.Host, error) {
	return uc.repo.GetHostInfo(ctx, verifyCode)
}
