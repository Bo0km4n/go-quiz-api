package v1

import "github.com/gin-gonic/gin"

func Create(ctx *gin.Context) {

	ctx.JSON(200, gin.H{"message": "quiz create"})
	return
}

func Index(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "quiz Index"})
	return
}
