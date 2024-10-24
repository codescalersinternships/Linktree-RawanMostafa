package controllers

import (
	"context"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/codescalersinternships/Linktree-RawanMostafa/db"
	"github.com/codescalersinternships/Linktree-RawanMostafa/helpers"
	"github.com/codescalersinternships/Linktree-RawanMostafa/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func getUserIDFromToken(c *gin.Context) (userID string) {

	tokenString := c.GetHeader("Authorization")

	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token not provided"})
		return
	}

	tokenParts := strings.Split(tokenString, " ")
	if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authentication token"})
		c.Abort()
		return
	}

	tokenString = tokenParts[1]

	claims, err := helpers.VerifyToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	userID = claims["user_id"].(string)
	return
}

var linkCollection *mongo.Collection = db.OpenCollection(db.Client, "link")

func AddLink(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var link models.Link
	if err := c.BindJSON(&link); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	link.UserID = getUserIDFromToken(c)
	link.LinkID = primitive.NewObjectID().Hex()

	_, insertErr := linkCollection.InsertOne(ctx, link)
	if insertErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while creating link"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "link added successfully"})
}

func EditLink(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	linkID := c.Param("link_id")
	log.Println(linkID)

	var input struct {
		Url string `json:"url"`
	}
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	log.Println("New URL:", input.Url)

	update := bson.M{"$set": bson.M{"url": input.Url}}
	result, updateErr := linkCollection.UpdateOne(ctx, bson.M{"linkid": linkID}, update)
	if updateErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating link"})
		return
	}

	if result.MatchedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Link not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Link updated successfully"})
}
