package database

var _ Connection = &connection{}

type connection struct {
}

type Connection interface {
}

func NewConnection() Connection {
	conn := connection{}
	return &conn
}
