package router

import (
	"github.com/gin-gonic/gin"
	"github/akalitt/go-errors-example/internal/model"
	"github/akalitt/go-errors-example/pkg/header"
	"net/http"
)

func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	g.Use(gin.Recovery())
	g.Use(header.NoCache)
	g.Use(header.Options)
	g.Use(header.Secure)
	g.Use(header.RequestId)
	// Middlewares.
	g.Use(mw...)
	//g.Use(jwt.Token)

	// 404 Handler.

	tagAPI := InitBookAPI(model.DB)

	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	tag := g.Group("/v1/tags")
	{
		tag.POST("", tagAPI.Create)
		tag.GET("/:id", tagAPI.GetByID)
	}

	return g

}
