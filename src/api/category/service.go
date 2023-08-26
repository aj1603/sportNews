package category

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func get(ctx *gin.Context) {
	results, err := get_()

	if err != nil {
		ctx.JSON(500, err.Error())
		return
	}

	ctx.JSON(200, results)
}

func get_by_category_id(ctx *gin.Context) {
	id := ctx.Param("id")
	int_id, _ := strconv.Atoi(id)

	res := get_by_category_id_(int_id)

	ctx.JSON(200, res)
}
