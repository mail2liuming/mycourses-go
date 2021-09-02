package registeration

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mail2liuming/mycourses/courses"
	"gorm.io/gorm"
)

var registerService RegisterationService

func RegisterationController(r *gin.RouterGroup, db *gorm.DB) {
	//init
	registerService = NewRegisterationService(courses.NewSqlliteRepository(db), NewRegisterationRepository(db))

	r.GET("list/:userid", GetRegisteredCourses)
	r.POST("register", RegisterCourse)
	r.POST("deregister", DeregisterCourse)
	r.POST("update", UpdateCourseRegisteration)
}

func GetRegisteredCourses(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("userid"), 10, 32)
	if err != nil {
		c.JSON(401, "bad request")
		return
	}
	courses, err := registerService.listRegisteredCourses(uint(userID))
	if err != nil {
		c.JSON(404, "Not found")
		return
	}
	c.JSON(200, courses)
}

func RegisterCourse(c *gin.Context) {
	var course CourseRegisterDTO
	c.BindJSON(&course)
	registerID, err := registerService.resgisterCourse(&course)
	if err != nil {
		c.JSON(401, "bad request")
		return
	}
	c.JSON(200, gin.H{"registerID": registerID})
}

func DeregisterCourse(c *gin.Context) {}

func UpdateCourseRegisteration(c *gin.Context) {}
