package route

import "github.com/gin-gonic/gin"

func SetUp() error {
	engine := gin.Default()
	route(engine)
	user(engine)
	sqlRoute(engine)
	err := engine.Run(":8080")
	return err
}