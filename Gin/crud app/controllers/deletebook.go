package controllers

import (
	"gin/crud_app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteBook(c *gin.Context) {
	var book models.Book
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	models.DB.Delete(&book)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
