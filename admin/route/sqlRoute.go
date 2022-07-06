package route

import "github.com/gin-gonic/gin"

func sqlRoute(engin *gin.Engine)  {
	sql := engin.Group("/sql")

	sql.GET("/Two",sqlTwo)
	sql.GET("/One", sqlOne)
	sql.GET("/full",getScoreFull)
	sql.GET("/Three",sqlThree)
	sql.GET("/Four",sqlFour)
	sql.GET("/fifth",sqlFifth)
	sql.GET("/seventh",sqlSeventh)
}