package users

import (
	req "sport_news/src/api/user/schemas"

	"github.com/gin-gonic/gin"
)

func Controller(router *gin.RouterGroup) {
	router.POST("/register", req.Validate_create, register)
	router.POST("/login", req.Validate_login, login)
	router.POST("/update/:id", req.Validate_update, update)
}
