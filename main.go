package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
	"os"
)

var database *gorm.DB

func hello(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusOK, map[string]string{
		"check": "ok",
	})
}

func user (c *gin.Context) {
	err := database.Create(&User{
		Username:  "test",
		Email: "obaldan@hotmail.com",
		UserRoleID: "admin",
		Password: "1234",
	}).Error
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{
			"check": err.Error(),
		})
		return
	}

	c.AbortWithStatusJSON(http.StatusOK, map[string]string{
		"check": "ok",
	})
}

func buildConnectionString() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		"3306",
		os.Getenv("DB_SCHEMA"),
	)
}

func main() {
	var err error
	router := gin.Default()

	router.GET("", hello)
	router.GET("/user", user)

	connectionString := buildConnectionString()
	fmt.Println(connectionString)
	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{
		//Logger:
	})
	if err != nil {
		panic(err)
	}

	database = db
	err = router.Run(":5000")
	if err != nil {
		panic(err)
	}
}
