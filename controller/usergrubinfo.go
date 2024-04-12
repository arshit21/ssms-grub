package controller

import (
	"context"
	"fmt"
	"net/http"
	"ssms_grub/database"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type RequestBody struct {
	S_ID string `json:"s_id"`
}

func GetUserGrubInfo(c *gin.Context) {
	var body RequestBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	s_id := body.S_ID
	fmt.Println("s_id is:", s_id)

	db, err := database.StartDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	grubsCollection := db.Collection("grubs")
	var grub database.Grub
	err = grubsCollection.FindOne(context.Background(), bson.M{"upcoming": true}).Decode(&grub)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "No upcoming grub found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	fmt.Println("Grub name: ", grub.Name)

	var userGrubInfo database.UserGrubInfo
	found := false
	for _, ugi := range grub.UserGrubInfo {
		if ugi.S_ID == s_id {
			userGrubInfo = ugi
			found = true
			break
		}
	}

	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}

	c.JSON(http.StatusOK, userGrubInfo)
}
