package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Blacklist struct {
	ID   int64
	Code string
}

type BlacklistRepo interface {
	GetBlacklistByCode(ctx context.Context, code string) (*Blacklist, bool, error)
	CreateBlacklist(ctx context.Context, code string) (*Blacklist, error)
	DeleteBlacklist(ctx context.Context, code string) error
	IsBlacklistNotFound(err error) bool
}

type BlacklistUseCase struct {
	repo BlacklistRepo
	log  *log.Helper
}

func NewBlacklistUseCase(repo BlacklistRepo, logger log.Logger) *BlacklistUseCase {
	return &BlacklistUseCase{repo: repo, log: log.NewHelper(log.With(logger, "module", "usecase/blacklist"))}
}

func (uc *BlacklistUseCase) CheckCodeBlacklisted(ctx context.Context, verifyCode string) (bool, error) {
	_, has, err := uc.repo.GetBlacklistByCode(ctx, verifyCode)
	if err != nil {
		return false, err
	}
	return has, nil
}

func (uc *BlacklistUseCase) BlockCode(ctx context.Context, verifyCode string) error {
	_, has, err := uc.repo.GetBlacklistByCode(ctx, verifyCode)
	if err != nil {
		return err
	}
	if has {
		return nil
	}

	_, err = uc.repo.CreateBlacklist(ctx, verifyCode)
	return err
}

func (uc *BlacklistUseCase) UnblockCode(ctx context.Context, verifyCode string) error {
	if err := uc.repo.DeleteBlacklist(ctx, verifyCode); err != nil && !uc.repo.IsBlacklistNotFound(err) {
		return err
	}
	return nil
}
