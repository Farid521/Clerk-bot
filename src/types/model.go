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