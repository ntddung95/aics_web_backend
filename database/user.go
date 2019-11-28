package database

import (
	"context"
	"fmt"
	"log"
	"errors"

	logger "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func init() {
	logger.Debug("Init mongodb")
	//clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	clientOptions := options.Client().ApplyURI("mongodb://mongo:27017")
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

func getUser(user string)(result UserLogin, err error){
        collection := client.Database("aics_web").Collection("user")
        logger.Debug("User: ", user)
        filter := bson.D{{"user", user}}
        err = collection.FindOne(context.TODO(), filter).Decode(&result)
        if err != nil {
                logger.Error(err.Error())
                return result, err
        }
	return result, nil
}
// LoginUser get user info from mongodb
func LoginUser(user, password string) (success bool, permission string) {
	logger.SetLevel(logger.DebugLevel)
	result, err := getUser(user)
	if err != nil {
		logger.Error(err.Error())
		return false, ""
	}
	logger.Debug("Result: ",result)
	if result.Password == password {
		return true, result.Permission
	}
	return false, ""

}

func UserRegister(user, password, permission string)(success bool, err error){
	_, err = getUser(user)
	if err == nil{
		return false, errors.New("User existed")
	}
	collection := client.Database("aics_web").Collection("user")
	new_user := bson.M{"user":user, "password":password, "permission":permission}
	_, err = collection.InsertOne(context.TODO(), new_user)
	if err != nil{
		return false, err
	}
	return true, nil
}

func UserChangePass(user, oldpass, newpass string) error {
	res_user, err := getUser(user)
	if err != nil{
		return errors.New("User not existed")
	}
	logger.Debug(res_user.Password)
	if oldpass != res_user.Password {
		return errors.New("Old pass not match")
	}
	collection := client.Database("aics_web").Collection("user")
	filter := bson.M{"user":user}
	update := bson.M{"$set": bson.M{"password":newpass}}
	_, err = collection.UpdateOne(context.TODO(),filter, update)
	if err != nil {
		return err
	}
	return nil
}
