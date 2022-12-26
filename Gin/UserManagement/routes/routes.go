package routes

import (
	"github/usermanage/handlers"
	"github/usermanage/middleware"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Use(middleware.SetUserStatus())
	router.GET("/", handlers.ShowIndexPage)

	userRoutes := router.Group("/u")
	{
		// register
		userRoutes.GET("/register", middleware.EnsureNotLoggedIn(), handlers.ShowRegistrationPage)
		userRoutes.POST("/register", middleware.EnsureNotLoggedIn(), handlers.Register)
		// login and logout
		userRoutes.GET("/login", middleware.EnsureNotLoggedIn(), handlers.ShowLoginPage)
		userRoutes.POST("/login", middleware.EnsureNotLoggedIn(), handlers.PerformLogin)
		userRoutes.GET("/logout", middleware.EnsureLoggedIn(), handlers.Logout)
	}

	articleRoutes := router.Group("/article")
	{
		// route from Part 1 of the tutorial
		articleRoutes.GET("/view/:article_id", handlers.GetArticle)

		articleRoutes.GET("/create", middleware.EnsureLoggedIn(), handlers.ShowArticleCreationPage)

		articleRoutes.POST("/create", middleware.EnsureLoggedIn(), handlers.CreateArticle)
	}
	router.Run(":3000")

}
