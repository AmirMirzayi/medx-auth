package services

import (
	"context"
	"errors"
	"log"
	"medx/auth/domain/auth"
	"medx/auth/models"
	"medx/auth/util"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type dbInterface interface {
	Find(auth.LoginRequestBody) (*models.User, error)
	Create(models.User) (interface{}, error)
}

type db struct {
	dbUrl string
}

var Mongodb dbInterface = &db{dbUrl: ""}

func InitDB() (dbInterface, error) {

	dbconn, err := util.LoadEnv()

	if err != nil {
		panic("failed to load environment")
	}

	var DB dbInterface = &db{dbUrl: dbconn}
	return DB, nil
}

func (db *db) Find(request auth.LoginRequestBody) (*models.User, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.dbUrl))

	if err != nil {
		return nil, errors.New("failed to connect database")
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	coll := client.Database("medx").Collection("users")

	filter := bson.D{
		{Key: "username", Value: request.UserName},
		{Key: "password", Value: request.Password},
	}

	var user models.User
	err = coll.FindOne(context.TODO(), filter).Decode(&user)

	if err != nil {
		return nil, err
	}
	return &user, nil

}

func (db *db) Create(user models.User) (interface{}, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.dbUrl))

	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	coll := client.Database("medx").Collection("users")

	result, err := coll.InsertOne(context.TODO(), user)

	if err != nil {
		return nil, err
	}
	return result, nil

}
