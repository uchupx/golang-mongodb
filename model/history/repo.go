package history

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type HistoryRepo interface {
	FindByUserId(ctx context.Context, id primitive.ObjectID) ([]History, error)
	Insert(ctx context.Context, history History) error
}
