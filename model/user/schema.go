package user

type User struct {
	Username string `bson:"Username" json:"username"`
	Password string `bson:"Password" json:"password"`
	Email    string `bson:"Email" json:"email"`
	Nama     string `bson:"Nama" json:"nama"`
}
