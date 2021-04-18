package history

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type historyModel struct {
	collection *mongo.Collection
}

func (m historyModel) FindByUserId(ctx context.Context, id primitive.ObjectID) ([]History, error) {
	var results []History
	findOptions := options.Find()
	idStr := id.Hex()

	raws, err := m.collection.Find(ctx, bson.M{"user_id": idStr}, findOptions)
	if err != nil {
		return nil, err
	}

	for raws.Next(ctx) {
		var elem History
		err := raws.Decode(&elem)
		if err != nil {
			return nil, err
		}

		results = append(results, elem)
	}

	if err := raws.Err(); err != nil {
		return nil, err
	}

	raws.Close(ctx)

	return results, nil
}

func (m historyModel) Insert(ctx context.Context, history History) error {
	insertResult, err := m.collection.InsertOne(ctx, history)
	if err != nil {
		return err
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	return nil
}

func NewHistoryModel(db *mongo.Database) *historyModel {
	var model historyModel

	collection := db.Collection("histories")
	model = historyModel{
		collection: collection,
	}

	return &model
}
