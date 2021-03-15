package history

type History struct {
	Deskripsi     string  `bson:"deskripsi" json:"deskripsi"`
	Jumlah        uint64  `bson:"jumlah" json:"jumlah"`
	Keterangan    *string `bson:"keterangan" json:"keterangan"`
	UsersUsername string  `bson:"users_username" json:"users_username"`
}
