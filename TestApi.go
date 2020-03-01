package main

import (
	"fmt"

	"./usersrepo"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	router, err := setupRouter()

	if err != nil {
		fmt.Println(err)
		return
	}

	router.Run(":8999")
}

func setupRouter() (*gin.Engine, error) {
	router := gin.Default()
	connectionString := "host=localhost port=5432 user=postgres dbname=TestGoDB password=Tp4tci2s4u2g! sslmode=disable"
	var db, err = gorm.Open("postgres", connectionString)

	//cant reach the db
	if err != nil {
		return router, err
	}

	//not matter what happens
	//we can close the connection
	defer db.Close()

	db.AutoMigrate(&usersrepo.User{})

	// Enable Logger, show detailed log
	db.LogMode(true)

	// This handler will match /user/john but will not match /user/ or /user
	router.GET("/user/:userId", usersrepo.GetUserFromDatabase(db))

	router.POST("/user", usersrepo.SaveUserToDatabase(db))
	return router, err
}
