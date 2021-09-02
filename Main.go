package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mail2liuming/mycourses/courses"
	"github.com/mail2liuming/mycourses/infrastructure"
	"github.com/mail2liuming/mycourses/registeration"
	"github.com/mail2liuming/mycourses/users"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&users.UserModel{})
	db.AutoMigrate(&courses.CourseModel{})
	db.AutoMigrate(registeration.RegisterationModel{})
}

func main() {

	db := infrastructure.Init()
	Migrate(db)

	r := gin.Default()

	v1 := r.Group("/api")
	users.UserController(v1.Group("/users"))
	courses.CoursesController(v1.Group("courses"))
	registeration.RegisterationController(v1.Group("registeration"), db)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
