package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/CircleConnectApp/post-service/database"
	"github.com/CircleConnectApp/post-service/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PostController struct {
	DB *mongo.Database
}

func NewPostController(db *mongo.Database) *PostController {
	return &PostController{DB: db}
}

func (pc *PostController) CreatePost(c *gin.Context) {
	log.Println("CreatePost: Endpoint hit")

	var req models.CreatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("CreatePost: Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Printf("CreatePost: Request body: %+v", req)

	userID, exists := c.Get("user_id")
	if !exists {
		log.Println("CreatePost: User not authenticated")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}
	log.Printf("CreatePost: User ID: %v", userID)

	now := time.Now()
	post := models.Post{
		UserID:      userID.(int),
		CommunityID: req.CommunityID,
		Title:       req.Title,
		Content:     req.Content,
		MediaURLs:   req.MediaURLs,
		Tags:        req.Tags,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	log.Printf("CreatePost: Post to insert: %+v", post)

	result, err := pc.DB.Collection(database.PostCollection).InsertOne(context.Background(), post)
	if err != nil {
		log.Printf("CreatePost: Error inserting post: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create post"})
		return
	}

	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		post.ID = oid
		log.Printf("CreatePost: Post created with ID: %v", post.ID)
	}

	c.JSON(http.StatusCreated, gin.H{"post": post})
}

func (pc *PostController) GetAllPosts(c *gin.Context) {
	ctx := context.Background()
	var posts []models.Post

	findOptions := options.Find()
	findOptions.SetSort(bson.D{{Key: "created_at", Value: -1}})

	cursor, err := pc.DB.Collection(database.PostCollection).Find(ctx, bson.M{}, findOptions)
	if err != nil {
		log.Printf("Error fetching posts: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch posts"})
		return
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &posts); err != nil {
		log.Printf("Error decoding posts: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode posts"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"posts": posts})
}

func (pc *PostController) GetPostByID(c *gin.Context) {
	id := c.Param("id")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	var post models.Post
	err = pc.DB.Collection(database.PostCollection).FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&post)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
			return
		}
		log.Printf("Error fetching post: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch post"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"post": post})
}

func (pc *PostController) GetPostsByCommunity(c *gin.Context) {
	communityID := c.Param("community_id")

	communityIDInt := 0
	_, err := fmt.Sscanf(communityID, "%d", &communityIDInt)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid community ID"})
		return
	}

	ctx := context.Background()
	var posts []models.Post

	findOptions := options.Find()
	findOptions.SetSort(bson.D{{Key: "created_at", Value: -1}})

	cursor, err := pc.DB.Collection(database.PostCollection).Find(ctx, bson.M{"community_id": communityIDInt}, findOptions)
	if err != nil {
		log.Printf("Error fetching posts: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch posts"})
		return
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &posts); err != nil {
		log.Printf("Error decoding posts: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode posts"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"posts": posts})
}

func (pc *PostController) GetPostsByUser(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	ctx := context.Background()
	var posts []models.Post

	findOptions := options.Find()
	findOptions.SetSort(bson.D{{Key: "created_at", Value: -1}})

	cursor, err := pc.DB.Collection(database.PostCollection).Find(ctx, bson.M{"user_id": userID.(int)}, findOptions)
	if err != nil {
		log.Printf("Error fetching posts: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch posts"})
		return
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &posts); err != nil {
		log.Printf("Error decoding posts: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode posts"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"posts": posts})
}
