package service

import (
	v1 "dipper/api/server/service/v1"
	"dipper/app/server/service/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

const (
	codeSuccess    = 200
	codeBadRequest = 400
	codeConflict   = 409
	codeFailed     = 500

	msgSuccess      = "success"
	msgDuplicated   = "duplicated"
	msgInvalidParam = "invalid parameter"
	msgFailed       = "failed"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewServerService)

type ServerService struct {
	v1.UnimplementedHostServer

	host      *biz.HostUseCase
	blacklist *biz.BlacklistUseCase
	log       *log.Helper
}

func NewServerService(host *biz.HostUseCase, blacklist *biz.BlacklistUseCase, logger log.Logger) *ServerService {
	return &ServerService{
		host:      host,
		blacklist: blacklist,
		log:       log.NewHelper(log.With(logger, "module", "service/server")),
	}
}
