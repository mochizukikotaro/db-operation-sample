package sample

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/mochizukikotaro/golang-orm-sample/model"
)

func SampleSql() {
	users := []model.User{}

	db, err := sql.Open("postgres",
		"user=postgres password=pw dbname=postgres sslmode=disable")
	rows, _ := db.Query(`SELECT * FROM "users"`)
	defer rows.Close()
	user := model.User{}
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
