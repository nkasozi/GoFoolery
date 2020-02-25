package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB
var err error

type user struct {
	gorm.Model
	Name string
}

func getUserFromDatabase(c *gin.Context) {

	//read the Id
	userID := c.Param("userId")

	//query the db for the user
	var aUser user
	db.Where("id = ?", userID).First(&aUser)

	//user wasnt found
	if aUser.ID == 0 {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	//we are good
	c.JSON(http.StatusOK, aUser)
	return
}

func saveUserToDatabase(c *gin.Context) {

	aUser := new(user)

	err := c.Bind(aUser)

	//cant parse request
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	//save user
	db.Create(&aUser)

	//we are good
	c.String(http.StatusOK, "SUCCESS")
	return
}

func main() {

	connectionString := "host=localhost port=5432 user=postgres dbname=TestGoDB password=Tp4tci2s4u2g! sslmode=disable"
	db, err = gorm.Open("postgres", connectionString)

	//cant reach the db
	if err != nil {
		fmt.Println(err)
		return
	}

	//not matter what happens
	//we can close the connection
	defer db.Close()

	db.AutoMigrate(&user{})

	// Enable Logger, show detailed log
	db.LogMode(true)

	router := gin.Default()

	// This handler will match /user/john but will not match /user/ or /user
	router.GET("/user/:userId", getUserFromDatabase)

	router.POST("/user", saveUserToDatabase)

	router.Run(":8999")
}
