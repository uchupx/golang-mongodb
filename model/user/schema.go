package user

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Username string             `bson:"username,omitempt" json:"username"`
	Password string             `bson:"password,omitempt" json:"password"`
	Email    string             `bson:"email,omitempt" json:"email"`
	Nama     string             `bson:"nama,omitempt" json:"nama"`
}
