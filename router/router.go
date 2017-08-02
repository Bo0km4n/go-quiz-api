package router

import "github.com/gin-gonic/gin"

func V1(app *gin.Engine) *gin.Engine {
	v2 := app.Group("/api/v1")
	{
		quiz_router := v1.Group("/quiz")
		quiz_router.POST("", quiz.Create)
		quiz_router.GET("", quiz.Index)
	}
	{
		ranking_router := v1.Group("/ranking")
		ranking_router.GET("", ranking.Index)
		ranking_router.PUT("", ranking.Update)
	}
}
