package biz

import (
	"context"
	"dipper/app/server/service/internal/conf"
	util "dipper/pkg/tls"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/holdno/snowFlakeByGo"
)

type Host struct {
	ID         int64
	Name       string
	Manager    string
	Phone      string
	VerifyCode string
	CertNum    int64
}

type HostCert struct {
	Num  int64
	Cert []byte
	Key  []byte
}

type HostRepo interface {
	CreateHost(ctx context.Context, h *Host) (*Host, error)
	DeleteHost(ctx context.Context, id int64) error
	ListHost(ctx context.Context) ([]*Host, error)
	UpdateHostWithoutVerifyCode(ctx context.Context, h *Host) (*Host, error)
	UpdateCertNumByHostID(ctx context.Context, hostID int64, certNum int64) (*Host, error)
	GetHostByVerifyCode(ctx context.Context, verifyCode string) (*Host, bool, error)
	GetHostByCertNum(ctx context.Context, certNum int64) (*Host, bool, error)
	GetHostByName(ctx context.Context, name string) (*Host, bool, error)
	IsHostNotFound(err error) bool
}

type HostUseCase struct {
	certFile string
	keyFile  string

	idWorker *snowFlakeByGo.Worker

	repo HostRepo
	log  *log.Helper
}

func NewHostUseCase(conf *conf.Server, repo HostRepo, logger log.Logger) *HostUseCase {
	uc := &HostUseCase{
		certFile: conf.Tls.CertFile,
		keyFile:  conf.Tls.KeyFile,

		repo: repo,
		log:  log.NewHelper(log.With(logger, "module", "usecase/host")),
	}

	var err error
	uc.idWorker, err = snowFlakeByGo.NewWorker(0)
	if err != nil {
		uc.log.Fatalf("failed to new id worker: %v", err)
	}
	return uc
}

func (uc *HostUseCase) Create(ctx context.Context, host *Host) (*Host, bool, error) {
	_, has, err := uc.repo.GetHostByName(ctx, host.Name)
	if err != nil {
		return nil, false, err
	}
	if has {
		return nil, true, nil
	}
	rv, err := uc.repo.CreateHost(ctx, host)
	if err != nil {
		return nil, false, err
	}
	return rv, false, nil
}

func (uc *HostUseCase) Delete(ctx context.Context, id int64) error {
	if err := uc.repo.DeleteHost(ctx, id); err != nil && !uc.repo.IsHostNotFound(err) {
		return err
	}
	return nil
}

func (uc *HostUseCase) List(ctx context.Context) ([]*Host, error) {
	return uc.repo.ListHost(ctx)
}

func (uc *HostUseCase) Update(ctx context.Context, host *Host) (*Host, error) {
	return uc.repo.UpdateHostWithoutVerifyCode(ctx, host)
}

func (uc *HostUseCase) GetHostByVerifyCode(ctx context.Context, verifyCode string) (*Host, bool, error) {
	return uc.repo.GetHostByVerifyCode(ctx, verifyCode)
}

func (uc *HostUseCase) GetHostByCertNum(ctx context.Context, certNum int64) (*Host, bool, error) {
	return uc.repo.GetHostByCertNum(ctx, certNum)
}

func (uc *HostUseCase) UpdateCertNumByHostID(ctx context.Context, hostID, certNum int64) (*Host, error) {
	return uc.repo.UpdateCertNumByHostID(ctx, hostID, certNum)
}

func (uc *HostUseCase) GenHostCert(_ context.Context) (*HostCert, error) {
	num := uc.idWorker.GetId()
	cert, key, err := util.GenClient(uc.certFile, uc.keyFile, num)
	if err != nil {
		return nil, err
	}
	return &HostCert{
		Num:  num,
		Cert: cert,
		Key:  key,
	}, nil
}
