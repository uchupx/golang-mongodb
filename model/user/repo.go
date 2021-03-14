package user

import "context"

type UserRepo interface {
	FindAll(ctx context.Context) ([]User, error)
	Insert(ctx context.Context, users []User) error
}
