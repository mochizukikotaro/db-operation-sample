package sample

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/mochizukikotaro/db-operation-sample/model"
)

func SampleSqlx() {
	users := model.CreateUsers()

	db, _ := sqlx.Open("postgres",
		"user=postgres password=pw dbname=postgres sslmode=disable")
	db.Select(&users, "select * from users")
	fmt.Println("sqlx:")
	fmt.Println(users)
}
