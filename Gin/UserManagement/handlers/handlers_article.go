package handlers

import (
	"github/usermanage/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func showIndexPage(c *gin.Context) {
	articles := models.GetAllArticles()

	// Call the render function with the name of the template to render
	models.Render(c, "index.html",
		gin.H{
			"title":   "Home Page",
			"payload": articles})
}
func GetArticle(c *gin.Context) {
	// Check if the article ID is valid
	if articleID, err := strconv.Atoi(c.Param("article_id")); err == nil {
		// Check if the article exists
		if article, err := models.GetArticleByID(articleID); err == nil {
			models.Render(c, "article.html", gin.H{
				"title":   article.Title,
				"payload": article,
			})

		} else {
			// If the article is not found, abort with an error
			c.AbortWithError(http.StatusNotFound, err)
		}

	} else {
		// If an invalid article ID is specified in the URL, abort with an error
		c.AbortWithStatus(http.StatusNotFound)
	}
}
func ShowIndexPage(c *gin.Context) {
	articles := models.GetAllArticles()
	models.Render(c, "index.html",
		gin.H{
			"title":   "Home Page",
			"payload": articles,
		})
}

func ShowArticleCreationPage(c *gin.Context) {
	models.Render(c, "create-article.html",
		gin.H{
			"title": "Create New Article",
		})
}

func CreateArticle(c *gin.Context) {
	title := c.PostForm("title")
	content := c.PostForm("content")
	if article, err := models.CreateNewArticle(title, content); err == nil {
		models.Render(c, "submission-successful.html",
			gin.H{
				"title":   "Submission Successful",
				"payload": article,
			})
	} else {
		c.AbortWithStatus(http.StatusBadRequest)
	}

}
