package app

import (
	_ "github.com/lib/pq"
	"github.com/midaef/emmet-server/app/endpoint"
	"github.com/midaef/emmet-server/app/endpoint/auth"
	"github.com/midaef/emmet-server/app/repository"
	"github.com/midaef/emmet-server/app/service"
	"github.com/midaef/emmet-server/config"
	"github.com/midaef/emmet-server/dependers/database"
	app_auth "github.com/midaef/emmet-server/extra/auth"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
	"time"
)

type App struct {
	logger *zap.Logger
	config *config.Config
}

func NewApp(logger *zap.Logger, config *config.Config) *App {
	return &App{
		logger: logger,
		config: config,
	}
}

func (app *App) StartApp(certPath string) error {
	startTime := time.Now().UnixNano()

	connectionAddr := &database.Connection{
		Host:     app.config.DB.Host,
		Port:     app.config.DB.Port,
		User:     app.config.DB.User,
		Password: app.config.DB.Password,
		DBName:   app.config.DB.DBName,
		SSLMode:  app.config.DB.SSLMode,
		CertPath: certPath,
	}

	connectionAddrStr := database.GenerateAddr(connectionAddr)

	db, err := database.NewDB(connectionAddrStr)
	if err != nil {
		return err
	}

	app.logger.Info(" successfully connected to database",
		zap.String("addr", app.config.DB.Host+":"+app.config.DB.Port),
		zap.String("db_name", app.config.DB.DBName),
		zap.String("user", app.config.DB.User),
		zap.Int64("duration", time.Now().UnixNano()-startTime),
	)

	store := repository.NewRepository(db)

	service := service.NewService(store, app.config)

	endpointContainer := app.InitEndpointContainer(service)

	listener, err := net.Listen("tcp", ":"+app.config.Server.Port)
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()

	app_auth.RegisterAuthServer(grpcServer, endpointContainer.AuthService)

	app.logger.Info("emmet-server successfully started",
		zap.String("addr", app.config.Server.IP+":"+app.config.Server.Port),
		zap.Int64("duration", time.Now().UnixNano()-startTime),
	)

	if err := grpcServer.Serve(listener); err != nil {
		return err
	}

	return nil
}

func (app *App) InitEndpointContainer(service *service.Service) *endpoint.EndpointContainer {
	authServices := auth.NewAuthEndpoint(service, app.config)

	serviceContainer := endpoint.NewEndpointContainer(
		authServices,
	)

	return serviceContainer
}
