package router

import (
	"social-media-app/controllers"
	"social-media-app/middleware"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func StartApp(DB *pgxpool.Pool) *gin.Engine {
	router := gin.Default()
	router.Use(func(c *gin.Context) {
		c.Set("DB", DB)
		c.Next()
	})

	commentRouter := router.Group("/v1/post/comment")
	{
		commentRouter.POST("/", middleware.CommentValidator(), controllers.CreateComment)
	}

	friendRouter := router.Group("/v1/friend")
	{
		friendRouter.POST("/", middleware.FriendValidator(), controllers.CreateFriend)
		friendRouter.GET("/", controllers.GetListOfFriend)
		friendRouter.DELETE("/", middleware.FriendValidator(), controllers.RemoveFriend)
	}

	userAccount := router.Group("/v1/user")
	{
		userAccount.POST("/register", middleware.RegisterValidator(), controllers.UserRegister)
		userAccount.POST("/login", middleware.AuthValidator(), controllers.UserLogin)
		userAccount.PATCH("/", middleware.UpdateAccountValidator(), controllers.UpdateAccountController)
		
		// Linking email and phone number User
		userAccount.POST("/link", middleware.Authentication(), middleware.LinkEmailValidator(), controllers.LinkEmail)
		userAccount.POST("/link/phone", middleware.Authentication(), middleware.LinkPhoneValidator(), controllers.LinkPhone)
	}
  
	postRouter := router.Group("v1/post")
	{
		postRouter.POST("/", middleware.AuthValidator(), middleware.PostValidator(), controllers.CreatePost)
		postRouter.GET("/", middleware.AuthValidator(), controllers.GetPost)
	}

	router.GET("/seed-test", controllers.CreateSeed)

	router.GET("/health-check", controllers.ServerCheck)

	return router
}
