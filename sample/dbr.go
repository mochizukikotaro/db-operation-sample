package sample

import (
	"fmt"
	"github.com/gocraft/dbr"
	_ "github.com/lib/pq"
	"github.com/mochizukikotaro/db-operation-sample/model"
)

func SampleDbr() {
	users := model.CreateUsers()

	conn, _ := dbr.Open("postgres",
		"user=postgres password=pw dbname=postgres sslmode=disable", nil)
	sess := conn.NewSession(nil)
	sess.Select("id", "name").From("users").Load(&users)
	fmt.Println("dbr")
	fmt.Println(users)
}
