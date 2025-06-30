package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Mahasiswa struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Nama       string             `bson:"nama" json:"nama"`
	NPM        int                `bson:"npm" json:"npm"`
	Prodi      string             `bson:"prodi" json:"prodi"`
	Fakultas   string             `bson:"fakultas" json:"fakultas"`
	Alamat     Alamat             `bson:"alamat" json:"alamat"`
	Minat      []string           `bson:"minat" json:"minat"`
	MataKuliah []MataKuliah       `bson:"mata_kuliah" json:"mata_kuliah"`
}

type Alamat struct {
	Jalan     string `bson:"jalan" json:"jalan"`
	Kelurahan string `bson:"kelurahan" json:"kelurahan"`
	Kota      string `bson:"kota" json:"kota"`
}

type MataKuliah struct {
	Kode  string `bson:"kode" json:"kode"`
	Nama  string `bson:"nama" json:"nama"`
	Nilai int    `bson:"nilai" json:"nilai"`
}

type UserLogin struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
	Role     string `json:"role" bson:"role"`
}

type Payload struct {
	User string `json:"user"`
	Role string `json:"role"`
}