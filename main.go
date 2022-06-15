package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/eibrahimarisoy/car_rental/pkg/config"
	db "github.com/eibrahimarisoy/car_rental/pkg/database"
	"github.com/gin-gonic/gin"
)

func main() {

	cfg, err := config.LoadConfig("./pkg/config/config-local")

	if err != nil {
		fmt.Println("Failed to load config:", err)
	}

	DB := db.NewPsqlDB(cfg)

	fmt.Println("DB:", DB)

	r := gin.Default()

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", cfg.ServerConfig.Port),
		Handler:      r,
		ReadTimeout:  time.Duration(cfg.ServerConfig.ReadTimeoutSecs) * time.Second,
		WriteTimeout: time.Duration(cfg.ServerConfig.WriteTimeoutSecs) * time.Second,
	}

	srv.ListenAndServe()
}
