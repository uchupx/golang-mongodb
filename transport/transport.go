package transport

import (
	"github.com/uchupx/golang-mongodb/config"
	"github.com/uchupx/golang-mongodb/model/user"
	"github.com/uchupx/golang-mongodb/transport/reqres"
	"go.mongodb.org/mongo-driver/mongo"
)

type TransportHandler struct {
	mongoConn *mongo.Database
	userRepo  user.UserRepo

	userRequest *reqres.UserRequest
}

func (t TransportHandler) NewUserRequest(conf *config.Config) *reqres.UserRequest {
	if t.userRequest == nil {
		userReq := reqres.UserRequest{
			UserRepo: t.newUserRepo(conf),
		}

		t.userRequest = &userReq
	}
	return t.userRequest
}

func (t TransportHandler) newUserRepo(conf *config.Config) user.UserRepo {
	if t.userRepo == nil {
		userConn := user.NewUserModel(t.NewMongoConn(conf))
		t.userRepo = userConn
	}

	return t.userRepo
}

func (t TransportHandler) NewMongoConn(conf *config.Config) *mongo.Database {
	if t.mongoConn == nil {
		db, err := config.ConnectionMongo(conf)
		if err != nil {
			panic(err)
		}

		t.mongoConn = db
	}

	return t.mongoConn
}
