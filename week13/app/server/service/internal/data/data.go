package data

import (
	"context"
	"dipper/app/server/service/internal/conf"
	"dipper/app/server/service/internal/data/ent"
	"dipper/app/server/service/internal/data/ent/migrate"
	"time"

	"github.com/allegro/bigcache/v3"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	_ "github.com/mattn/go-sqlite3"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewEntClient, NewHostRepo, NewBlacklistRepo)

// Data .
type Data struct {
	db    *ent.Client
	cache *bigcache.BigCache
	log   *log.Helper
}

func NewEntClient(conf *conf.Data, logger log.Logger) *ent.Client {
	log := log.NewHelper(log.With(logger, "module", "server-service/data/ent"))

	client, err := ent.Open(
		conf.Database.Driver,
		conf.Database.Source,
	)
	if err != nil {
		log.Fatalf("failed opening connection to db: %v", err)
	}
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background(), migrate.WithForeignKeys(false)); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return client
}

// NewData .
func NewData(entClient *ent.Client, logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper(log.With(logger, "module", "server-service/data"))

	cache, err := bigcache.NewBigCache(bigcache.DefaultConfig(time.Hour * 365))
	if err != nil {
		return nil, nil, err
	}

	d := &Data{
		db:    entClient,
		cache: cache,
		log:   log,
	}
	return d, func() {
		if err := d.db.Close(); err != nil {
			log.Error(err)
		}
		if err := d.cache.Close(); err != nil {
			log.Error(err)
		}
	}, nil
}

func isDBNotFound(err error) bool {
	if err == nil {
		return false
	}
	return ent.IsNotFound(err)
}
