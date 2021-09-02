package users

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func UserController(r *gin.RouterGroup) {
	r.POST("/register", UserRegisteration)
	r.POST("/login", UserLogin)
}

func UserRegisteration(c *gin.Context) {
	fmt.Println("User register!")
}

func UserLogin(c *gin.Context) {
	fmt.Println("User login!")
}
