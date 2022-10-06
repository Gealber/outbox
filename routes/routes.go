package routes

import (
	controllerCat "github.com/Gealber/outbox/controllers/cat"
	repositoryCat "github.com/Gealber/outbox/repositories/cat"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// InitRoutes initialize routes on gin engine.
func InitRoutes(
	g *gin.Engine,
	dbApp *gorm.DB,
) {
	// repositories.
	catStore := repositoryCat.New(dbApp)

	// controllers.
	catCont := controllerCat.New(catStore)

	api := g.Group("/api/v1")
	{
		RoutesCat(api, catCont)
	}
}

// RoutesCat definition of gin routes for cat handlers.
func RoutesCat(g *gin.RouterGroup, catsController *controllerCat.Controller) {
	g.POST("/cats", catsController.Create)
	g.PATCH("/cats/:id", catsController.Update)
}
