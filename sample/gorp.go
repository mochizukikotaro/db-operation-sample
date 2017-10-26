package sample

import (
	"database/sql"
	"fmt"
	"github.com/go-gorp/gorp"
	_ "github.com/lib/pq"
	"github.com/mochizukikotaro/golang-orm-sample/model"
)

func SampleGorp() {
	users := model.CreateUsers()
	dbmap := initDb()
	defer dbmap.Db.Close()
	dbmap.Select(&users, "select * from users")
	fmt.Println("gorp")
	fmt.Println(users)
}

func initDb() *gorp.DbMap {
	db, _ := sql.Open("postgres",
		"user=postgres password=pw dbname=postgres sslmode=disable")
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}
	dbmap.AddTableWithName(model.CreateUser(), "users").SetKeys(true, "Id")
	return dbmap
}
