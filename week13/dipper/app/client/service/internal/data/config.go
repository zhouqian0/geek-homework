package data

import (
	"context"
	"dipper/app/client/service/internal/biz"
	"dipper/app/client/service/internal/data/ent"
	"dipper/app/client/service/internal/data/ent/cfg"
	"github.com/go-kratos/kratos/v2/log"
)

const (
	recordUndeleted = iota
	recordDeleted
)

type ConfigRepo struct {
	data *Data
	log  *log.Helper
}

func NewConfigRepo(data *Data, logger log.Logger) biz.ConfigRepo {
	return &ConfigRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/config")),
	}
}

func (r *ConfigRepo) BatchCreatConfig(ctx context.Context, cfgs ...*biz.Config) ([]*biz.Config, error) {
	bulk := make([]*ent.CfgCreate, len(cfgs))
	for i, c := range cfgs {
		bulk[i] = r.data.db.Cfg.
			Create().
			SetKey(c.Key).
			SetValue(c.Value)
	}
	_, err := r.data.db.Cfg.
		CreateBulk(bulk...).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return cfgs, nil
}

func (r *ConfigRepo) CreatConfig(ctx context.Context, c *biz.Config) (*biz.Config, error) {
	_, err := r.data.db.Cfg.
		Create().
		SetKey(c.Key).
		SetValue(c.Value).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *ConfigRepo) DeleteConfig(ctx context.Context, key string) error {
	return r.data.db.Cfg.
		Update().
		SetIsDeleted(recordDeleted).
		Where(cfg.KeyEQ(key)).
		Exec(ctx)
}

func (r *ConfigRepo) BatchDeleteConfig(ctx context.Context, keys ...string) error {
	return r.data.db.Cfg.
		Update().
		SetIsDeleted(recordDeleted).
		Where(cfg.KeyIn(keys...)).
		Exec(ctx)
}

func (r *ConfigRepo) UpdateConfig(ctx context.Context, c *biz.Config) (*biz.Config, error) {
	if err := r.data.db.Cfg.
		Update().
		SetValue(c.Value).
		Where(cfg.KeyEQ(c.Key), cfg.IsDeletedEQ(recordUndeleted)).
		Exec(ctx); err != nil {
		return nil, err
	}
	return c, nil
}

func (r *ConfigRepo) BatchUpdateConfig(ctx context.Context, cfgs ...*biz.Config) ([]*biz.Config, error) {
	batchUpdateConfig := func(ctx context.Context, tx *ent.Tx) error {
		for _, c := range cfgs {
			if err := tx.Cfg.
				Update().
				SetValue(c.Value).
				Where(cfg.KeyEQ(c.Key), cfg.IsDeletedEQ(recordUndeleted)).
				Exec(ctx); err != nil {
				return err
			}
		}
		return nil
	}
	if err := withTx(ctx, r.data.db, batchUpdateConfig); err != nil {
		return nil, err
	}
	return cfgs, nil
}

func (r *ConfigRepo) GetConfig(ctx context.Context, key string) (*biz.Config, bool, error) {
	po, err := r.data.db.Cfg.
		Query().
		Where(cfg.KeyEQ(key), cfg.IsDeletedEQ(recordUndeleted)).
		Only(ctx)
	if err != nil {
		if isDBNotFound(err) {
			return nil, false, nil
		}
		return nil, false, err
	}
	return &biz.Config{
		Key:   po.Key,
		Value: po.Value,
	}, true, nil
}

func (r *ConfigRepo) BatchGetConfig(ctx context.Context, keys ...string) ([]*biz.Config, error) {
	pos, err := r.data.db.Cfg.
		Query().
		Where(cfg.KeyIn(keys...), cfg.IsDeletedEQ(recordUndeleted)).
		All(ctx)
	if err != nil {
		return nil, err
	}

	rvs := make([]*biz.Config, len(pos))
	for i, po := range pos {
		rvs[i] = &biz.Config{
			Key:   po.Key,
			Value: po.Value,
		}
	}
	return rvs, nil
}

func (r *ConfigRepo) UpdateCache(key string, value []byte) error {
	return r.data.cache.Set(key, value)
}

func (r *ConfigRepo) GetCache(key string) ([]byte, bool, error) {
	val, err := r.data.cache.Get(key)
	if err != nil {
		if isCacheNotFound(err) {
			return nil, false, nil
		}
		return nil, false, err
	}
	return val, true, nil
}

func (r *ConfigRepo) DeleteCache(key string) error {
	return r.data.cache.Delete(key)
}

func (r *ConfigRepo) IsDBNotFound(err error) bool {
	return isDBNotFound(err)
}

func (r *ConfigRepo) IsCacheNotFound(err error) bool {
	return isCacheNotFound(err)
}
