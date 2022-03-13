package service

import (
	v1 "dipper/api/client/service/v1"
	"dipper/app/client/service/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

const (
	codeSuccess      = 200
	codeUnauthorized = 401
	codeFailed       = 500

	msgSuccess      = "success"
	msgUnauthorized = "unauthorized"
	msgFailed       = "failed"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewClientService)

type ClientService struct {
	v1.UnimplementedClientServer

	machine *biz.MachineUseCase
	config  *biz.ConfigUseCase
	host    *biz.HostUseCase
	log     *log.Helper
}

func NewClientService(machine *biz.MachineUseCase, config *biz.ConfigUseCase, host *biz.HostUseCase, logger log.Logger) *ClientService {
	return &ClientService{
		machine: machine,
		config:  config,
		host:    host,
		log:     log.NewHelper(log.With(logger, "module", "service/client")),
	}
}
