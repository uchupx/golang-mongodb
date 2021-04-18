package user

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type userModel struct {
	collection *mongo.Collection
}

func (m userModel) FindAll(ctx context.Context) ([]User, error) {
	var results []User
	findOptions := options.Find()

	usersRaw, err := m.collection.Find(ctx, bson.D{{}}, findOptions)
	if err != nil {
		return nil, err
	}

	for usersRaw.Next(ctx) {
		var elem User
		err := usersRaw.Decode(&elem)
		if err != nil {
			return nil, err
		}

		results = append(results, elem)
	}

	if err := usersRaw.Err(); err != nil {
		return nil, err
	}

	usersRaw.Close(ctx)

	return results, nil
}

func (m userModel) Insert(ctx context.Context, user User) error {
	count, err := m.collection.CountDocuments(ctx, bson.M{"username": user.Username})
	if err != nil {
		return err
	}

	if count == 0 {
		insertResult, err := m.collection.InsertOne(ctx, user)
		if err != nil {
			return err
		}

		fmt.Println("Inserted a single document: ", insertResult.InsertedID)
		return nil
	} else {
		return fmt.Errorf("Username is exist")
	}
}

func NewUserModel(db *mongo.Database) *userModel {
	var model userModel

	collection := db.Collection("users")
	model = userModel{
		collection: collection,
	}

	return &model
}
