package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/eibrahimarisoy/car_rental/pkg/config"
	db "github.com/eibrahimarisoy/car_rental/pkg/database"
	graceful "github.com/eibrahimarisoy/car_rental/pkg/graceful"
	router "github.com/eibrahimarisoy/car_rental/pkg/router"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	// swagger embed files
)

// @title           Car Rental API
// @version         1.0
// @description     This is a Car Rental API implementation.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  eibrahimarisoy@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth
func main() {

	cfg, err := config.LoadConfig("./pkg/config/config-local")

	if err != nil {
		fmt.Println("Failed to load config:", err)
	}

	DB := db.NewPsqlDB(cfg)

	r := gin.Default()

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", cfg.ServerConfig.Port),
		Handler:      r,
		ReadTimeout:  time.Duration(cfg.ServerConfig.ReadTimeoutSecs) * time.Second,
		WriteTimeout: time.Duration(cfg.ServerConfig.WriteTimeoutSecs) * time.Second,
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	rootRouter := r.Group(cfg.ServerConfig.RoutePrefix)
	router.InitiliazeRoutes(rootRouter, DB, cfg)

	rootRouter.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, nil)
	})

	rootRouter.GET("/readyz", func(c *gin.Context) {
		db, err := DB.DB()
		if err != nil {
			// zap.L().Fatal("cannot get sql database instance", zap.Error(err))
			fmt.Println("cannot get sql database instance", err)
		}
		if err := db.Ping(); err != nil {
			// zap.L().Fatal("cannot ping database", zap.Error(err))
			fmt.Println("cannot ping database", err)
		}
		c.JSON(http.StatusOK, nil)
	})

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Listen error: %v", err)
		}
	}()
	log.Println("Car Rental API service started")

	// Wait for interrupt signal to gracefully shutdown the server with
	graceful.ShutdownGin(srv, time.Duration(cfg.ServerConfig.TimeoutSecs*int64(time.Second)))
}
