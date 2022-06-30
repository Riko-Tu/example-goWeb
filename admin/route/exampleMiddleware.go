package route

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// say中间件方式1
func sayNextMiddle(c *gin.Context) {
	say, ok := c.GetQuery("say")
	if !ok{
		say = "hi"
	}
	start := time.Now()
	c.Set("say",say)
	c.Next()
	cost := time.Since(start)
	fmt.Printf("%v\n",cost)

}

// 中间件方式2
func jsonAbortMiddle() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, ok := c.GetQuery("name")
		if !ok{
			// 驳回请求
			c.Abort()
			c.JSON(http.StatusBadRequest,gin.H{"err":"name不存在"})
		}else {
			//执行后续函数
			c.Next()
		}
	}
}