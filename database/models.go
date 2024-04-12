package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type menu struct {
	Vegetarian    []string `json:"vegetarian" bson:"vegetarian"`
	NonVegetarian []string `json:"nonvegetarian" bson:"nonvegetarian"`
}

type price struct {
	VegetarianPrice    int `json:"vegetarianprice" bson:"vegetarianprice"`
	NonVegeterianPrice int `json:"nonvegetarianprice" bson:"nonvegetarianprice"`
}

type Grub struct {
	ID           primitive.ObjectID `json:"_id" bson:"_id"`
	Name         string             `json:"name" bson:"name"`
	Logo         string             `json:"logo" bson:"logo"`
	Menu         menu               `json:"menu" bson:"menu"`
	Day          string             `json:"day" bson:"day"`
	Date         time.Time          `json:"date" bson:"date"`
	Price        price              `json:"price" bson:"price"`
	UserGrubInfo []UserGrubInfo     `json:"usergrubinfo" bson:"usergrubinfo"`
	Upcoming     bool               `json:"upcoming" bson:"upcoming"`
}

type UserGrubInfo struct {
	S_ID           string `json:"s_id" bson:"s_id"`
	VegSigning     bool   `json:"vegsigning" bson:"vegsigning"`
	NonVegSigning  bool   `json:"nonvegsigning" bson:"nonvegsigning"`
	VegScanned     bool   `json:"vegscanned" bson:"vegscanned"`
	NonVegScanned  bool   `json:"nonvegscanned" bson:"nonvegscanned"`
	Volunteering   bool   `json:"volunteering" bson:"volunteering"`
	InternalMember bool   `json:"internalmember" bson:"internalmember"`
}

func InsertGrub(db *mongo.Database, grub Grub) (*mongo.InsertOneResult, error) {
	collection := db.Collection("grubs")
	res, err := collection.InsertOne(context.Background(), grub)
	return res, err
}
