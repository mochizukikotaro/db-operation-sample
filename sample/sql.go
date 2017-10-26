package sample

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/mochizukikotaro/db-operation-sample/model"
)

func SampleSql() {
	users := model.CreateUsers()

	db, err := sql.Open("postgres",
		"user=postgres password=pw dbname=postgres sslmode=disable")
	rows, _ := db.Query(`SELECT * FROM "users"`)
	defer rows.Close()
	user := model.CreateUser()
	for rows.Next() {
		err = rows.Scan(&user.Id, &user.Name)
		if err != nil {
			fmt.Println(err)
		}
		users = append(users, user)
	}
	fmt.Println("sql:")
	fmt.Println(users)
}
