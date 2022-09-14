package app

import (
	"context"
	"fmt"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/wiryawanyurih/ecommerce-api/internal/config"
	"github.com/wiryawanyurih/ecommerce-api/internal/handler"
	"github.com/wiryawanyurih/ecommerce-api/internal/repository"
	"github.com/wiryawanyurih/ecommerce-api/internal/service"
	"github.com/wiryawanyurih/ecommerce-api/pkg/database/mongodb"
	"github.com/wiryawanyurih/ecommerce-api/pkg/database/redis"
)

func Run(path string) {
	log.Info("application startup...")
	log.Info("logger initialized")

	cfg := config.GetConfig(path)
	log.Info("config created")

	mongoClient, err := mongodb.NewClient(context.Background(),
		cfg.DB.URI, cfg.DB.Username, cfg.DB.Password)
	if err != nil {
		log.Fatal(err)
	}
	log.Info("mongo is connected")
	db := mongoClient.Database(cfg.DB.Database)

	redisClient, err := redis.NewClient(cfg)
	if err != nil {
		log.Fatal(err)
	}
	log.Info("redis is connected")

	repos := repository.NewRepositories(db)
	services := service.NewServices(service.Deps{
		Repos:       repos,
		RedisClient: redisClient,
	})
	handlers := handler.NewHandler(services, cfg)
	log.Info("services, repositories and handlers initialized")

	server := &http.Server{
		Handler:      handlers.Init(),
		Addr:         fmt.Sprintf(":%s", cfg.Listen.Port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Infof("server started on port %s", cfg.Listen.Port)

	log.Fatal(server.ListenAndServe())
}
