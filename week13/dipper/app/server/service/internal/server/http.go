package server

import (
	v1 "dipper/api/server/service/v1"
	"dipper/app/server/service/internal/conf"
	"dipper/app/server/service/internal/service"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/swagger-api/openapiv2"
	"github.com/gorilla/handlers"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"

	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, logger log.Logger, tp *tracesdk.TracerProvider, s *service.ServerService) *http.Server {
	// todo
	//log := log.NewHelper(log.With(logger, "server", "server-service/server/http"))
	//
	//tlsConf, err := pkgTLSConfig(c.Tls.CaCertFile, c.Tls.CertFile, c.Tls.KeyFile)
	//if err != nil {
	//	log.Fatalf("failed package tls config: %v", err)
	//}
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			tracing.Server(
				tracing.WithTracerProvider(tp)),
			logging.Server(logger),
		),
		http.Filter(handlers.CORS(
			handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
			handlers.AllowedOrigins([]string{"*"}),
		)),
		//http.TLSConfig(tlsConf),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	v1.RegisterHostHTTPServer(srv, s)

	// todo
	openAPIhandler := openapiv2.NewHandler()
	srv.HandlePrefix("/q/", openAPIhandler)
	return srv
}
