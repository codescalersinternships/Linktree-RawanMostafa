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

func Login(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var user models.User

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
