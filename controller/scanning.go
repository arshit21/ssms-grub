package controller

import (
	"context"
	"net/http"
	"ssms_grub/database"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type ScanRequestBody struct {
	S_ID          string `json:"s_id"`
	VegScanned    string `json:"vegscanned"`
	NonVegScanned string `json:"nonvegscanned"`
}

func HandleScan(c *gin.Context) {
	var body ScanRequestBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	vegScanned, err := strconv.ParseBool(body.VegScanned)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid value for vegscanned"})
		return
	}

	nonVegScanned, err := strconv.ParseBool(body.NonVegScanned)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid value for nonvegscanned"})
		return
	}

	db, err := database.StartDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	grubsCollection := db.Collection("grubs")
	var grub database.Grub
	err = grubsCollection.FindOne(context.Background(), bson.M{"upcoming": true}).Decode(&grub)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var updatedUserGrubInfo database.UserGrubInfo

	for i, ugi := range grub.UserGrubInfo {
		if ugi.S_ID == body.S_ID {
			if ugi.VegScanned || ugi.NonVegScanned {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Already scanned"})
				return
			}

			if vegScanned {
				if ugi.VegSigning {
					grub.UserGrubInfo[i].VegScanned = true
				} else {
					c.JSON(http.StatusOK, gin.H{"message": "Veg signing not done"})
					return
				}
			}

			if nonVegScanned {
				if ugi.NonVegSigning {
					grub.UserGrubInfo[i].NonVegScanned = true
				} else {
					c.JSON(http.StatusOK, gin.H{"message": "Non-veg signing not done"})
					return
				}
			}

			updatedUserGrubInfo = grub.UserGrubInfo[i]
			break
		}
	}

	_, err = grubsCollection.UpdateOne(context.Background(), bson.M{"_id": grub.ID}, bson.M{"$set": bson.M{"usergrubinfo": grub.UserGrubInfo}})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedUserGrubInfo)
}
