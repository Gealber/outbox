package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/Gealber/outbox/config"
	"github.com/Gealber/outbox/database"
	"github.com/Gealber/outbox/middlewares"
	"github.com/Gealber/outbox/msgrelay"
	repositoryEvent "github.com/Gealber/outbox/repositories/event"
	"github.com/Gealber/outbox/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	ctx := context.Background()
	cfg := config.Config()

	// setup databse.
	dbApp, err := database.New(ctx, cfg)
	if err != nil {
		log.Fatal(err)
	}

	eventRepo := repositoryEvent.New(dbApp)
	pubsubClient, err := msgrelay.NewPubSub(ctx, cfg.GCP.ProjectID)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		err := eventRepo.ChangeFeed(ctx, pubsubClient)
		if err != nil {
			log.Fatal(err)
		}
	}()

	gin.SetMode(cfg.Gin.Mode)
	g := gin.Default()

	g.Use(middlewares.Cors(cfg.Cors.Origins))

	setUpNotFoundRoute(g)
	routes.InitRoutes(g, dbApp)

	srvAddr := fmt.Sprintf(":%s", cfg.App.Port)
	if err := g.Run(srvAddr); err != nil {
		panic(err)
	}
}

func setUpNotFoundRoute(r *gin.Engine) {
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": fmt.Sprint("not found"),
		})
	})
}
