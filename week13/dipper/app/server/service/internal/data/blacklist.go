package data

import (
	"context"
	"dipper/app/server/service/internal/biz"
	"dipper/app/server/service/internal/data/ent"
	"dipper/app/server/service/internal/data/ent/blacklist"
	"github.com/go-kratos/kratos/v2/log"
)

const (
	recordUndeleted = iota
	recordDeleted
)

type BlacklistRepo struct {
	data *Data
	log  *log.Helper
}

func NewBlacklistRepo(data *Data, logger log.Logger) biz.BlacklistRepo {
	return &BlacklistRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/blacklist")),
	}
}

func (r BlacklistRepo) GetBlacklistByCode(ctx context.Context, code string) (*biz.Blacklist, bool, error) {
	po, err := r.data.db.Blacklist.
		Query().
		Where(blacklist.CodeEQ(code), blacklist.IsDeletedEQ(recordUndeleted)).
		Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, false, nil
		}
		return nil, false, err
	}
	return &biz.Blacklist{
		ID:   po.ID,
		Code: po.Code,
	}, true, nil
}

func (r BlacklistRepo) CreateBlacklist(ctx context.Context, code string) (*biz.Blacklist, error) {
	po, err := r.data.db.Blacklist.
		Create().
		SetCode(code).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.Blacklist{
		ID:   po.ID,
		Code: po.Code,
	}, nil
}

func (r BlacklistRepo) DeleteBlacklist(ctx context.Context, code string) error {
	return r.data.db.Blacklist.
		Update().
		SetIsDeleted(recordDeleted).
		Where(blacklist.CodeEQ(code)).
		Exec(ctx)
}

func (r BlacklistRepo) IsBlacklistNotFound(err error) bool {
	return isDBNotFound(err)
}
