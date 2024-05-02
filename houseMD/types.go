package housemd

type TUser struct {
	Id            int64
	UserLogin     string
	UserFirstName string
	MsgCount      int
}

type TMessage struct {
	ID   int
	User TUser
	Text string
}
