package courses

import "github.com/gin-gonic/gin"

func CoursesController(r *gin.RouterGroup) {
	r.POST("/create", CreateCourse)
	r.POST("/update", UpdateCourse)
	r.POST("/delete", DeteleCourse)
}

func CreateCourse(c *gin.Context) {}

func UpdateCourse(c *gin.Context) {}

func DeteleCourse(c *gin.Context) {}
