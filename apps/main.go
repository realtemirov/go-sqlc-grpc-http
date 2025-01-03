package apps

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/realtemirov/go-sqlc-grpc-http/config"
	db "github.com/realtemirov/go-sqlc-grpc-http/db/sqlc"

	"github.com/realtemirov/go-sqlc-grpc-http/pkg/logger"
	gprcservices "github.com/realtemirov/go-sqlc-grpc-http/transport/gprc_services"
	"github.com/realtemirov/go-sqlc-grpc-http/transport/gprc_services/clients"
	"github.com/realtemirov/go-sqlc-grpc-http/transport/handlers"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type App struct {
	cfg *config.Config
	db  db.Store
}

func (a *App) Run() {
	log.Println("Initializing...")

	cfg := config.Load()
	a.cfg = cfg

	logger.SetLogger(&cfg.Logging)

	zap.L().Info("Initializing done...")

	log.Println("Initializing configs done...")

	var (
		store db.Store
	)

	ctx, cancel := context.WithTimeout(context.Background(), cfg.Project.Timeout)
	defer cancel()

	store = db.NewStore(ctx, cfg.PSQL.URI)
	a.db = store

	grpcClient := clients.NewGrpcClients(cfg)
	grpcServer := gprcservices.New(gprcservices.GrpcServerParams{
		Cfg:     cfg,
		DB:      store,
		Clients: grpcClient,
	})

	lis, err := net.Listen("tcp", cfg.Grpc.URL)
	if err != nil {
		panic(err)
	}

	go func() {
		zap.L().Info("starting grpc server on " + cfg.Grpc.URL)

		err = grpcServer.Serve(lis)
		if err != nil {
			panic(err)
		}
	}()

	router := a.setUpHTTP()
	go func() {
		zap.L().Info("starting http server on " + cfg.HTTP.URL)

		err = router.Run(cfg.HTTP.URL)
		if err != nil {
			panic(err)
		}
	}()

	zap.L().Info("All services are up and running")
	_ = zap.L().Sync()
	a.gracefulShutdown(ctx, grpcServer, cancel)
}

func (a *App) setUpHTTP() *gin.Engine {
	router := gin.Default()

	switch a.cfg.Mode {
	case string(config.DEVELOPMENT):
		gin.SetMode(gin.DebugMode)
	default:
		gin.SetMode(gin.ReleaseMode)
	}

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = append(config.AllowHeaders, "*")

	handler := handlers.NewHandler(a.cfg, a.db)

	router.Use(cors.New(config))
	v1 := router.Group("/v1")

	gwMux := handlers.New(context.Background(), a.cfg)

	router.Static("/swagger", "./doc/swagger")
	v1.Any("/*any", func(c *gin.Context) {
		gwMux.ServeHTTP(c.Writer, c.Request)
	})

	router.GET("health", handler.HealthCheck)
	return router
}

func (a *App) gracefulShutdown(ctx context.Context, grpcServer *grpc.Server, cancel context.CancelFunc) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c

	var wg sync.WaitGroup

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		zap.L().Info("shutting down")

		grpcServer.GracefulStop()

		zap.L().Info("shutdown successfully called")
		wg.Done()
	}(&wg)

	go func() {
		wg.Wait()
		cancel()
	}()

	<-ctx.Done()
	os.Exit(0)
}
