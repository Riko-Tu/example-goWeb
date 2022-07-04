package route

import "github.com/gin-gonic/gin"

func sqlRoute(engin *gin.Engine)  {
	sql := engin.Group("/sql")

	sql.GET("/sqlTwo",sqlTwo)
	sql.GET("/sqlOne", sqlOne)
}