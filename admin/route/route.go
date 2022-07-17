package route

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "turan/example-goWeb/admin/docs"
)

//路由组
func user(engin *gin.Engine)  {
	engin.POST("/sendEmailCode",sendEmailCode)
	engin.POST("/emailLogin",emailLogin)
	engin.POST("/registerEmail",RegisterEmail)
}

func SetUp() error {
	engine := gin.Default()
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	user(engine)
	err := engine.Run(":8080")
	return err
}


