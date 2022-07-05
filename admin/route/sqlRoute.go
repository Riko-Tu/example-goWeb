package route

import "github.com/gin-gonic/gin"

func sqlRoute(engin *gin.Engine)  {
	sql := engin.Group("/sql")

	sql.GET("/sqlTwo",sqlTwo)
	sql.GET("/sqlOne", sqlOne)
	sql.GET("/full",getScoreFull)
	sql.GET("/sqlThree",sqlThree)
	sql.GET("/sqlFour",sqlFour)
}