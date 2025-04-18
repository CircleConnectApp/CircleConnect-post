package routes

import (
	"github.com/CircleConnectApp/post-service/controllers"
	"github.com/CircleConnectApp/post-service/middleware"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)


func SetupRoutes(r *gin.Engine, db *mongo.Database) {
	
	postController := controllers.NewPostController(db)

	
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"service": "post-service",
		})
	})

	
	api := r.Group("/api")

	
	api.GET("/posts", postController.GetAllPosts)
	api.GET("/posts/:id", postController.GetPostByID)
	api.GET("/communities/:community_id/posts", postController.GetPostsByCommunity)

	
	auth := api.Group("/")
	auth.Use(middleware.AuthMiddleware())
	{
		
		auth.POST("/posts", postController.CreatePost)
		auth.GET("/user/posts", postController.GetPostsByUser)
	}
}
