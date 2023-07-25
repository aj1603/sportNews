package category

import (
	"github.com/gin-gonic/gin"
)

func Controller(router *gin.RouterGroup) {
	router.GET("/all", get)
	router.GET("/get-by-category-id/:id", get_by_category_id)
}
