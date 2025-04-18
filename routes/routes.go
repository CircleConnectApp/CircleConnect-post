package routes

import (
	"log"

	"github.com/CircleConnectApp/post-service/controllers"
	"github.com/CircleConnectApp/post-service/middleware"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRoutes(r *gin.Engine, db *mongo.Database) {
	log.Println("Setting up routes...")

	postController := controllers.NewPostController(db)

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"service": "post-service",
		})
	})

	api := r.Group("/api")
	log.Println("Registering /api routes...")

	api.GET("/posts", func(c *gin.Context) {
		log.Println("GET /api/posts route hit")
		postController.GetAllPosts(c)
	})
	api.GET("/posts/:id", func(c *gin.Context) {
		log.Println("GET /api/posts/:id route hit")
		postController.GetPostByID(c)
	})
	api.GET("/communities/:community_id/posts", func(c *gin.Context) {
		log.Println("GET /api/communities/:community_id/posts route hit")
		postController.GetPostsByCommunity(c)
	})

	// Authenticated routes
	auth := api.Group("/")
	auth.Use(func(c *gin.Context) {
		log.Println("AuthMiddleware applied")
		middleware.AuthMiddleware()(c)
	})
	{
		auth.POST("/posts", func(c *gin.Context) {
			log.Println("POST /api/posts route hit")
			postController.CreatePost(c)
		})
		auth.GET("/user/posts", func(c *gin.Context) {
			log.Println("GET /api/user/posts route hit")
			postController.GetPostsByUser(c)
		})
	}
	log.Println("Routes registered successfully.")
}
