package middlewares

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	_defaultOrigin = "*"
)

func Cors(origins []string) gin.HandlerFunc {
	if len(origins) == 0 {
		origins = []string{_defaultOrigin}
	}

	middleware := cors.New(cors.Config{
		AllowOrigins: origins,
		AllowMethods: []string{"POST", "OPTIONS", "GET", "PUT", "PATCH"},
		AllowHeaders: []string{
			"Origin",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"X-CSRF-Token",
			"Authorization",
			"Accept",
			"Cache-Control",
			"X-Requested-With",
		},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	})

	return middleware
}
