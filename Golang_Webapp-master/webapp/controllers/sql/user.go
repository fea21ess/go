package sql

type User struct {
	Id   int
	Name string `orm:"size(100)"`
	Password string `orm:"size(100)"`
	Email string `orm:"size(100)"`
}