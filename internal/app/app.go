package app

import (
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"

	"github.com/midaef/emmet-server/configs"
	"github.com/midaef/emmet-server/internal/pkg"
	"github.com/midaef/emmet-server/internal/repository"
	"github.com/midaef/emmet-server/internal/service"
	"github.com/midaef/emmet-server/pkg/helpers"
	"github.com/midaef/emmet-server/internal/api"
)

func Run(config *configs.Config) {
	logger, err := pkg.ConfigureLogger(config.Logger.LogLevel)
	if err != nil {
		logger.Error("Configure logger error", zap.Error(err))
	}

	conn := pkg.NewConnection(config.Database.Uri)
	err = conn.Open()
	if err != nil {
		logger.Error("Configure connection error", zap.Error(err))
	}

	defer conn.DB.Close()

	logger.Debug("Connected to administration database")

	repositories := repository.NewRepositories(conn.DB)

	hasher := helpers.NewHasher(config.Token.Salt)

	jwtManager, err := helpers.NewJWT(config.Token.SecretKey)
	if err != nil {
		logger.Error("JWT manager error", zap.Error(err))
	}

	deps := &service.Dependencies{
		Repository: repositories,
		Hasher:     hasher,
		JWTManager: jwtManager,
	}
	services := service.NewServices(deps)

	s := grpc.NewServer()
	api.RegisterAuthServer(s, services.AuthService)
	api.RegisterUserServer(s, services.UserService)

	logger.Info("Started emmet-server",
		zap.String("host", config.Server.Host),
		zap.String("port", config.Server.Port))

	l, err := net.Listen("tcp", config.Server.Port)
	if err != nil {
		logger.Error("listen tcp error", zap.Error(err))
	}

	err = s.Serve(l)
	if err != nil {
		logger.Error("serve error", zap.Error(err))
	}
}