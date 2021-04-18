package history

type History struct {
	Deskripsi  string  `bson:"deskripsi" json:"deskripsi"`
	Jumlah     uint64  `bson:"jumlah" json:"jumlah"`
	Keterangan *string `bson:"keterangan" json:"keterangan"`
	UserId     string  `bson:"user_id" json:"user_id"`
}
