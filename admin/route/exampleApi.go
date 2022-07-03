package route

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
)


type data struct {
	Name string `json:"name"` //tag结构体返回小写字段
	Msg string `json:"msg"`
}




func route(engin *gin.Engine)  {
	engin.GET("/hello",sayNextMiddle,ginH)
	engin.GET("/json",jsonAbortMiddle(),structJson)
	engin.GET("/name",urlQuery)
	engin.POST("/login",formLogin)
	// tips: 避免路由匹配冲突
	engin.GET("/user/:num/name",urlPathQuery)
	// 该绑定在query或者form表单,rawJSon格式都适用
	engin.POST("/dataBind",structBind)
	engin.POST("/upload",uploadFile)

	//重定向
	engin.GET("/redirect", func(c *gin.Context) {
			c.Request.URL.Path = "/hello" //修改参数路径
			engin.HandleContext(c)		 //执行后续操作
	})

	//空路由
	engin.NoRoute(func(c *gin.Context) {
		c.Request.URL.Path = "/hello"
		engin.HandleContext(c)
	})



	sql := engin.Group("/sql")
	sql.GET("/course", sqlOne)
}


//路由组
func user(engin *gin.Engine)  {
	user := engin.Group("/user")
	user.GET("/hello",ginH)
}

// 自定义gin.h json格式
func ginH(c *gin.Context)  {
	//从中间件获取值
	say, _ := c.Get("say")
	c.JSON(http.StatusOK,
		gin.H{
		"msg":say,
		"code":http.StatusOK,
		"err":"xxx"})
}

// 返回结构体json
func structJson(c *gin.Context)  {
	data := data{
		"小明",
		"你好",
	}
	c.JSON(http.StatusOK,data)
}

//获取url ? 拼接的参数
func urlQuery(c *gin.Context)  {
	// 获取参数
	name ,ok:= c.GetQuery("name")
	if !ok{
		c.JSON(http.StatusBadRequest,gin.H{"err":"name不能为空"})
		return
	}
	age , ok := c.GetQuery("age")
	if !ok{
		c.JSON(http.StatusBadRequest,gin.H{"err":"age不能为空"})
		return
	}
	//c.DefaultQuery("name","小明")

	// 返回参数
	c.JSON(http.StatusOK,
		gin.H{
		"name":name,
		"age":age,
	})
}


// 接受form表单数据
func formLogin(c *gin.Context)  {
	// 接收数据
	username, ok := c.GetPostForm("username")
	if !ok{
		c.JSON(http.StatusBadRequest,gin.H{"err":"username不能为空"})
		return
	}
	password, ok := c.GetPostForm("password")
	if !ok{
		c.JSON(http.StatusBadRequest,gin.H{"err":"password不能为空"})
		return
	}

	// 返回
	data := gin.H{"username":username,"password":password}
	c.JSON(http.StatusOK,data)
}

// 获取url路劲参数
func urlPathQuery(c *gin.Context)  {
	num := c.Param("num")
	if num ==""{
		c.JSON(http.StatusBadRequest,gin.H{"err":"path参数有误"})
		return
	}
	c.JSON(http.StatusOK,gin.H{"num":num})


}

type Bind struct {
	Name string `form:"name" json:"name"`
	Age string  `form:"age" json:"age"`
}

// data参数绑定：适用三种提交方式
func structBind(c *gin.Context)  {
	var d Bind
	//反射
	err := c.ShouldBind(&d)
	if err != nil{
		c.JSON(http.StatusBadRequest,gin.H{"err":err.Error()})
		return
	}

	//返回
	c.JSON(http.StatusOK,d)
}

// 获取文件
func uploadFile(c *gin.Context)  {
	//读取文件
	file, err := c.FormFile("file")
	if err !=nil{
		c.JSON(http.StatusOK,gin.H{"err":err.Error()})
		return
	}

	//保存文件
	dst :=path.Join("./fileFold",file.Filename) //项目路径下
	err = c.SaveUploadedFile(file, dst)
	if err !=nil {
		c.JSON(http.StatusOK,gin.H{"err":err.Error()})
		return
	}

	//返回响应
	c.JSON(http.StatusOK,gin.H{"msg":"上传成功"})
}




