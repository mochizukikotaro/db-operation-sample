package sample

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/mochizukikotaro/golang-orm-sample/model"
)

func SampleGorm() {
	users := []model.User{}

	db, _ := gorm.Open("postgres",
		"user=postgres password=pw dbname=postgres sslmode=disable")
	defer db.Close()
	db.Find(&users) // SELECT * FROM users;
	fmt.Println("gorm")
	fmt.Println(users)
}
