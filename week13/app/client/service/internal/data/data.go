package data

import (
	"context"
	"crypto/tls"
	serverv1 "dipper/api/server/service/v1"
	"dipper/app/client/service/internal/conf"
	"dipper/app/client/service/internal/data/ent"
	"dipper/app/client/service/internal/data/ent/migrate"
	util "dipper/pkg/tls"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"time"

	"github.com/allegro/bigcache/v3"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewEntClient, NewConfigRepo, NewHostRepo, NewHostServer)

type hostServer struct {
	serverAddr string
	caCertFile string
	client     serverv1.HostClient
}

// Data .
type Data struct {
	db         *ent.Client
	cache      *bigcache.BigCache
	hostServer *hostServer
	log        *log.Helper
}

func NewEntClient(conf *conf.Data, logger log.Logger) *ent.Client {
	log := log.NewHelper(log.With(logger, "module", "client-service/data/ent"))

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
func NewData(conf *conf.Server, entClient *ent.Client, hostServer *hostServer, logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper(log.With(logger, "module", "client-service/data"))

	cache, err := bigcache.NewBigCache(bigcache.DefaultConfig(time.Hour * 365))
	if err != nil {
		return nil, nil, err
	}

	d := &Data{
		db:         entClient,
		cache:      cache,
		hostServer: hostServer,
		log:        log,
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

func withTx(ctx context.Context, client *ent.Client, fn func(ctx context.Context, tx *ent.Tx) error) error {
	tx, err := client.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if v := recover(); v != nil {
			_ = tx.Rollback()
			panic(v)
		}
	}()

	if err = fn(ctx, tx); err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = errors.Wrapf(err, "rolling back transaction: %v", rerr)
		}
		return err
	}
	if err = tx.Commit(); err != nil {
		return errors.Wrapf(err, "committing transaction: %v", err)
	}
	return nil
}

func NewHostServer(c *conf.Server) *hostServer {
	certPool, err := util.GenCertPool(c.Tls.CaCertFile)
	if err != nil {
		panic(err)
	}
	tlsConf := &tls.Config{RootCAs: certPool, InsecureSkipVerify: true}
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(c.Remote.Addr),
		grpc.WithMiddleware(recovery.Recovery()),
		grpc.WithTLSConfig(tlsConf),
	)
	if err != nil {
		panic(err)
	}
	return &hostServer{
		serverAddr: c.Remote.Addr,
		caCertFile: c.Tls.CaCertFile,
		client:     serverv1.NewHostClient(conn),
	}
}

func updateHostServer(server *hostServer, certPem, keyPem []byte) error {
	certPool, err := util.GenCertPool(server.caCertFile)
	if err != nil {
		return err
	}
	cert, err := tls.X509KeyPair(certPem, keyPem)
	if err != nil {
		return err
	}
	tlsConf := &tls.Config{RootCAs: certPool, Certificates: []tls.Certificate{cert}, InsecureSkipVerify: true}
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(server.serverAddr),
		grpc.WithMiddleware(recovery.Recovery()),
		grpc.WithTLSConfig(tlsConf),
	)
	if err != nil {
		return err
	}
	server.client = serverv1.NewHostClient(conn)
	return nil
}

func isCacheNotFound(err error) bool {
	if err == nil {
		return false
	}
	return errors.Is(err, bigcache.ErrEntryNotFound)
}

func isDBNotFound(err error) bool {
	if err == nil {
		return false
	}
	return ent.IsNotFound(err)
}
