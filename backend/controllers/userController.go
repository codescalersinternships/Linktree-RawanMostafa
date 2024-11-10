package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/codescalersinternships/Linktree-RawanMostafa/db"
	"github.com/codescalersinternships/Linktree-RawanMostafa/helpers"
	"github.com/codescalersinternships/Linktree-RawanMostafa/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection = db.OpenCollection(db.Client, "user")

func HashPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Fatal(err.Error())
	}
	return string(hashedPassword)
}

func VerifyPassword(userPassword string, providedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(providedPassword))
	return err == nil
}

// Signup		godoc
// @Summary		User signup
// @Description Register the user if its username doesn't exist and save it in the db
// @Tags		User
// @Accept		json
// @Produce		json
// @Param		user	body		models.User	true	"User Data"
// @Success		201		{object}	models.MsgResponse
// @Failure		400		{object}	models.ErrorResponse
// @Failure		500		{object}	models.ErrorResponse
// @Router		/user/register [post]
func Signup(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userNameCount, err := userCollection.CountDocuments(ctx, bson.M{"username": user.Username})
	if err != nil || userNameCount > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
		return
	}

	user.Password = HashPassword(user.Password)

	user.ID = primitive.NewObjectID().Hex()

	_, insertErr := userCollection.InsertOne(ctx, user)
	if insertErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while creating user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

// Login		godoc
// @Summary		User login
// @Description Logs a user is if credentials are valid and returns a JWT token
// @Tags		User
// @Accept		json
// @Produce		json
// @Param		loginRequest	body		models.LoginRequest	true	"Login Credentials"
// @Success		200				{object}	models.LoginResponse
// @Failure		400				{object}	models.ErrorResponse
// @Failure		401				{object}	models.ErrorResponse
// @Failure		500				{object}	models.ErrorResponse
// @Router		/user/login [post]
func Login(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var user models.LoginRequest

	var foundUser models.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := userCollection.FindOne(ctx, bson.M{"username": user.Username}).Decode(&foundUser)
	if err == mongo.ErrNoDocuments {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "username doesn't exist"})
		return
	}

	passwordIsValid := VerifyPassword(foundUser.Password, user.Password)
	if !passwordIsValid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid password"})
		return
	}

	token, tokenErr := helpers.GenerateToken(foundUser.ID)
	if tokenErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while generating the token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})

}

// UpdateBio godoc
// @Summary     Update user bio
// @Description Updates the bio for a specific user
// @Tags        User
// @Accept      json
// @Produce     json
// @Param       username path     string true "Username"
// @Param       bio      body     string true "User Bio"
// @Success		200				{object}	models.MsgResponse
// @Failure		400				{object}	models.ErrorResponse
// @Failure		404				{object}	models.ErrorResponse
// @Failure		500				{object}	models.ErrorResponse
// @Router		/user/update-bio/:username [put]
func UpdateBio(c *gin.Context) {
    var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
    defer cancel()

    username := c.Param("username")
    var updateData struct {
        Bio string `json:"bio" binding:"required"`
    }

    if err := c.BindJSON(&updateData); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Bio is required"})
        return
    }
	update := bson.M{"$set": bson.M{
		"bio":updateData.Bio,
	}}
    result, updateErr := userCollection.UpdateOne(ctx, bson.M{"username": username}, update)
    if updateErr != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating bio"})
        return
    }
	if result.MatchedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

    c.JSON(http.StatusOK, gin.H{"message": "Bio updated successfully"})
}
