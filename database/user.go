package database

import (
	"context"
	"fmt"
	"log"

	logger "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func init() {
	logger.Debug("Init mongodb")
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, _ = mongo.Connect(context.TODO(), clientOptions)

	err := client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

}

// UserLogin UserLogin
type UserLogin struct {
	User       string `json:"user"`
	Password   string `json:"pass"`
	Permission string `json:"permission"`
}

// LoginUser get user info from mongodb
func LoginUser(user, password string) (success bool, permission string) {
	logger.SetLevel(logger.DebugLevel)
	var result UserLogin
	collection := client.Database("aics_web").Collection("user")
	logger.Debug("User: ", user)
	filter := bson.D{{"user", user}}
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		logger.Error(err.Error())
		return false, ""
	}
	if result.Password == password {
		return true, result.Permission
	}
	return false, ""

}
