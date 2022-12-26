package handlers

import (
	"github/snippet/models"

	"github.com/gin-gonic/gin"
)

func ShowIndexPage(c *gin.Context) {
	articles := models.GetAllArticles()
	models.Render(c, "index.html",
		gin.H{
			"title":   "Home Page",
			"payload": articles,
		})
}
