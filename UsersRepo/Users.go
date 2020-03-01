package usersrepo

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//User => contains user details
type User struct {
	gorm.Model
	Name string
}

//SaveUserToDatabase => saves a given user to the database
func SaveUserToDatabase(db *gorm.DB) gin.HandlerFunc {

	fn := func(c *gin.Context) {

		aUser := new(User)

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
	return gin.HandlerFunc(fn)
}

//GetUserFromDatabase = > retrieves user from the db using the userId
func GetUserFromDatabase(db *gorm.DB) gin.HandlerFunc {

	fn := func(c *gin.Context) {

		//read the Id
		userID := c.Param("userId")

		//query the db for the user
		var aUser User
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
	return gin.HandlerFunc(fn)
}
