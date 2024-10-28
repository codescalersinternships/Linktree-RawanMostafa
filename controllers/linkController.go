package controllers

import (
	"context"
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

// AddLink		godoc
// @Summary		Add new link
// @Description Add a new link for the authenticated user
// @Tags		Link
// @Accept		json
// @Produce		json
// @Param		linkRequest 	body		models.LinkRequest	true	"Link Details"
// @Success		201				{object}	models.LinkResponse
// @Failure		400				{object}	models.ErrorResponse
// @Failure		500				{object}	models.ErrorResponse
// @Router		/api/v1/link [post]
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

	c.JSON(http.StatusCreated, gin.H{
		"message": "link added successfully",
		"linkid":  link.LinkID,
	})
}

// EditLink		godoc
// @Summary		Edit a link
// @Description Edit an existing link that's created by the authenticated user
// @Tags		Link
// @Accept		json
// @Produce		json
// @Param		linkRequest 	body		models.LinkRequest	true	"Link Details"
// @Success		200				{object}	models.MsgResponse
// @Failure		400				{object}	models.ErrorResponse
// @Failure		404				{object}	models.ErrorResponse
// @Failure		500				{object}	models.ErrorResponse
// @Router		/api/v1/link/:link_id [put]
func EditLink(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	linkID := c.Param("link_id")

	var input models.LinkRequest
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	update := bson.M{"$set": bson.M{
		"url":      input.Url,
		"platform": input.Platform,
	}}
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

// DeleteLink	godoc
// @Summary		Delete a link
// @Description Delete an existing link that's created by the authenticated user
// @Tags		Link
// @Accept		json
// @Produce		json
// @Success		200				{object}	models.MsgResponse
// @Failure		404				{object}	models.ErrorResponse
// @Failure		500				{object}	models.ErrorResponse
// @Router		/api/v1/link/:link_id [delete]
func DeleteLink(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	linkID := c.Param("link_id")

	result, deleteErr := linkCollection.DeleteOne(ctx, bson.M{"linkid": linkID})
	if deleteErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting link"})
		return
	}

	if result.DeletedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Link not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Link deleted successfully"})
}

// GetUserLinks	godoc
// @Summary		Get links
// @Description Get links of a specific username
// @Tags		Link
// @Accept		json
// @Produce		json
// @Success		200				{object}	models.MsgResponse
// @Failure		404				{object}	models.ErrorResponse
// @Failure		500				{object}	models.ErrorResponse
// @Router		/api/v1/link/:username [get]
func GetUserLinks(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	username := c.Param("username")
	var user models.User
	err := userCollection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Not user found with this username"})
		return
	}

	cursor, err := linkCollection.Find(ctx, bson.M{"userid": user.ID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error finding user links"})
		return
	}
	defer cursor.Close(ctx)

	var links []models.Link

	for cursor.Next(ctx) {
		var link models.Link
		if err := cursor.Decode(&link); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error decoding link"})
			return
		}
		links = append(links, link)
	}

	if len(links) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No links found for this user"})
		return
	}

	c.JSON(http.StatusOK, links)
}
