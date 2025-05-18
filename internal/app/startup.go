package app

import (
	"context"
	_ "go-template/docs"
	"go-template/internal/controller"
	"go-template/internal/repository"
	"go-template/internal/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/fx"
)

var Startup = fx.Options(
	fx.Provide(
		RegisterRoute,
		RegisterServer,
		// controller
		controller.NewUserController,
		// service
		service.NewUserService,
		// repository
		repository.NewUserRepository,
	),
	fx.Invoke(RegisterHooks),
)

func RegisterRoute(userController *controller.UserController) *gin.Engine {
	r := gin.New()

	// middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// route
	publicRoute := r.Group("public/api/v1")
	{
		publicRoute.GET("/user", userController.GetUser)
	}

	// health
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}

func RegisterServer(engine *gin.Engine) *http.Server {
	return &http.Server{
		Addr:    ":8080",
		Handler: engine,
	}
}

func RegisterHooks(lc fx.Lifecycle, server *http.Server) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				log.Println("Server starting on", server.Addr)
				if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					log.Fatalf("ListenAndServe error: %v", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Println("Shutting down server...")
			return server.Shutdown(ctx)
		},
	})
}