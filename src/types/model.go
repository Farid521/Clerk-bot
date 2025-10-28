package types

type Msg struct {
	Type string `bson:"msg-type"`
	MsgContent string `bson:"msg-content"`
}

type User struct {
	UserId     string `bson:"user-id"`
	UserName   string `bson:"user-name"`
	GlobalName string `bson:"global-name"`
}

type UserMsg struct {
	User User `bson:"user"`
	Msg  Msg  `bson:"msg"`
}

// JADWAL KULIAH TYPE
type JadwalKuliah struct {
	Matkul string `bson:"matkul"`
	Kode string	`bson:"kode"`
	Dosen string `bson:"dosen"`
	Hari string	`bson:"hari"`
	Waktu string `bson:"waktu"`
	Gedung string `bson:"gedung"`
	Ruangan string `bson:"ruangan"`
}