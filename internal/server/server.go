package server

import (
	"go.uber.org/zap"

	"github.com/midaef/emmet-server/internal/service"
)

type GRPCServer struct {
	services *service.Services
	logger   *zap.Logger
}

func NewGRPCServer(services *service.Services, logger *zap.Logger) *GRPCServer {
	return &GRPCServer{
		services: services,
		logger: logger,
	}
}