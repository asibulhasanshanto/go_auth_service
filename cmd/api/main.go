package main

import (
	"context"
	"net/http"

	"github.com/asibulhasanshanto/go_api/internal/api"
	"github.com/asibulhasanshanto/go_api/internal/api/handlers"
	"github.com/asibulhasanshanto/go_api/internal/config"
	"github.com/asibulhasanshanto/go_api/internal/conn"
	"github.com/asibulhasanshanto/go_api/internal/core"
	"github.com/asibulhasanshanto/go_api/internal/store"
	"github.com/asibulhasanshanto/go_api/pkg"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func main() {
	fx.New(
		fx.Provide(
			pkg.CustomLogger,
			config.LoadConfig,

			conn.ConnectPostgres,

			// handlers
			handlers.NewAuthHandler,

			// cores
			core.NewAuth,

			// stores
			store.NewUserStore,

			GinHttpServer,
			api.SetupRoutes,
		),
		fx.Invoke(func(r *gin.RouterGroup, log *zap.Logger) {
			log.Info("Setting up routes")
			log.Info("Starting the application")
		}),
	).Run()
}

func GinHttpServer(lc fx.Lifecycle, log *zap.Logger) *gin.Engine {
	r := gin.Default()

	srv := &http.Server{
		Addr:    ":" + "8082",
		Handler: r,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					log.Fatal("listen", zap.Error(err))
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})

	return r
}
