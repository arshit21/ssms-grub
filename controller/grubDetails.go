package controller

import (
	"context"
	"net/http"

	"ssms_grub/database"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateGrub(c *gin.Context) {
	var grub database.Grub
	if err := c.ShouldBindJSON(&grub); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db, err := database.StartDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	usersCollection := db.Collection("users")
	cursor, err := usersCollection.Find(context.Background(), bson.D{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer cursor.Close(context.Background())

	var users []struct {
		S_ID string `bson:"s_id"`
	}
	if err = cursor.All(context.Background(), &users); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	grub.ID = primitive.NewObjectID()

	for _, user := range users {
		grub.UserGrubInfo = append(grub.UserGrubInfo, database.UserGrubInfo{
			S_ID:           user.S_ID,
			VegSigning:     false,
			NonVegSigning:  false,
			VegScanned:     false,
			NonVegScanned:  false,
			Volunteering:   false,
			InternalMember: false,
		})
	}

	res, err := database.InsertGrub(db, grub)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"insertedID": res.InsertedID})
}

func GetAllGrubs(c *gin.Context) {
	db, err := database.StartDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	collection := db.Collection("grubs")
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer cursor.Close(context.Background())

	var grubs []database.Grub
	if err = cursor.All(context.Background(), &grubs); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, grubs)
}

func GetGrubByName(c *gin.Context) {
	db, err := database.StartDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	name := c.Param("name")

	collection := db.Collection("grubs")
	var grub database.Grub
	err = collection.FindOne(context.Background(), bson.M{"name": name}).Decode(&grub)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, grub)
}
