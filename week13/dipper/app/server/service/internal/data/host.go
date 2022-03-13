package data

import (
	"context"
	"dipper/app/server/service/internal/biz"
	"dipper/app/server/service/internal/data/ent"
	"dipper/app/server/service/internal/data/ent/host"
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

func (r *HostRepo) CreateHost(ctx context.Context, h *biz.Host) (*biz.Host, error) {
	po, err := r.data.db.Host.
		Create().
		SetName(h.Name).
		SetManager(h.Manager).
		SetPhone(h.Phone).
		SetVerifyCode(h.VerifyCode).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	h.ID = po.ID
	return h, nil
}

func (r *HostRepo) DeleteHost(ctx context.Context, id int64) error {
	var deleted uint8 = 1
	return r.data.db.Host.
		UpdateOneID(id).
		SetIsDeleted(deleted).
		Exec(ctx)
}

func (r *HostRepo) ListHost(ctx context.Context) ([]*biz.Host, error) {
	var undeleted uint8 = 0
	pos, err := r.data.db.Host.
		Query().
		Where(host.IsDeletedEQ(undeleted)).
		All(ctx)
	if err != nil {
		return nil, err
	}
	rvs := make([]*biz.Host, len(pos))
	for i, po := range pos {
		rvs[i] = &biz.Host{
			ID:         po.ID,
			Name:       po.Name,
			Manager:    po.Manager,
			Phone:      po.Phone,
			VerifyCode: po.VerifyCode,
			CertNum:    po.CertNum,
		}
	}
	return rvs, err
}

func (r *HostRepo) UpdateHostWithoutVerifyCode(ctx context.Context, h *biz.Host) (*biz.Host, error) {
	_, err := r.data.db.Host.
		UpdateOneID(h.ID).
		SetName(h.Name).
		SetManager(h.Manager).
		SetPhone(h.Phone).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return h, nil
}

func (r *HostRepo) GetHostByName(ctx context.Context, name string) (*biz.Host, bool, error) {
	var undeleted uint8 = 0
	po, err := r.data.db.Host.
		Query().
		Where(host.NameEQ(name), host.IsDeletedEQ(undeleted)).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, false, nil
		}
		return nil, false, err
	}
	return &biz.Host{
		ID:         po.ID,
		Name:       po.Name,
		Manager:    po.Manager,
		Phone:      po.Phone,
		VerifyCode: po.VerifyCode,
	}, true, nil
}

func (r *HostRepo) GetHostByVerifyCode(ctx context.Context, verifyCode string) (*biz.Host, bool, error) {
	var undeleted uint8 = 0
	po, err := r.data.db.Host.
		Query().
		Where(host.VerifyCodeEQ(verifyCode), host.IsDeletedEQ(undeleted)).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, false, nil
		}
		return nil, false, err
	}
	return &biz.Host{
		ID:         po.ID,
		Name:       po.Name,
		Manager:    po.Manager,
		Phone:      po.Phone,
		VerifyCode: po.VerifyCode,
	}, true, nil
}

func (r *HostRepo) GetHostByCertNum(ctx context.Context, certNum int64) (*biz.Host, bool, error) {
	var undeleted uint8 = 0
	po, err := r.data.db.Host.
		Query().
		Where(host.CertNumEQ(certNum), host.IsDeletedEQ(undeleted)).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, false, nil
		}
		return nil, false, err
	}
	return &biz.Host{
		ID:         po.ID,
		Name:       po.Name,
		Manager:    po.Manager,
		Phone:      po.Phone,
		VerifyCode: po.VerifyCode,
	}, true, nil
}

func (r *HostRepo) UpdateCertNumByHostID(ctx context.Context, hostID int64, certNum int64) (*biz.Host, error) {
	po, err := r.data.db.Host.
		UpdateOneID(hostID).
		SetCertNum(certNum).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.Host{
		ID:         po.ID,
		Name:       po.Name,
		Manager:    po.Manager,
		Phone:      po.Phone,
		VerifyCode: po.VerifyCode,
	}, nil
}

func (r *HostRepo) IsHostNotFound(err error) bool {
	if err == nil {
		return false
	}
	return ent.IsNotFound(err)
}
