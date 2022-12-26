package routes

import (
	"github/snippet/handlers"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	// Handle the index route
	router.GET("/", handlers.ShowIndexPage)
	router.GET("/:article_id", handlers.GetArticle)
	router.Run(":3001")
}
