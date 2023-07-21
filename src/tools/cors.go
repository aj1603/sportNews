package tools

import "github.com/gin-gonic/gin"

func Cors(ctx *gin.Context) {
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	ctx.Writer.Header().Set(
		"Access-Control-Allow-Headers",
		"Content-Type, "+
			"Content-Length, "+
			"Accept-Encoding, "+
			"X-CSRF-Token, "+
			"Authorization, "+
			"accept, "+
			"origin, "+
			"Cache-Control, "+
			"X-Requested-With",
	)
	ctx.Writer.Header().Set(
		"Access-Control-Allow-Methods",
		"OPTIONS, GET, POST, PUT, DELETE",
	)

	if ctx.Request.Method == "OPTIONS" {
		ctx.AbortWithStatus(204)
		return
	}

	ctx.Next()
}
