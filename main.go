package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

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
	ticker := time.NewTicker(30 * time.Second)
	done := make(chan bool)

	// spinning up message relay.
	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				err := msgrelay.Poll(ctx, eventRepo)
				if err != nil {
					fmt.Println("ERR: ", err)
				}
			}
		}
	}()

	gin.SetMode(cfg.Gin.Mode)
	g := gin.Default()

	g.Use(middlewares.Cors(cfg.Cors.Origins))

	setUpNotFoundRoute(g)
	routes.InitRoutes(g, dbApp)

	srvAddr := fmt.Sprintf(":%s", cfg.App.Port)
	if err := g.Run(srvAddr); err != nil {
		done <- true
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
