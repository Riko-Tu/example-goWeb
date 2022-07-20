package route

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"turan/example-goWeb/admin/cache"
	"turan/example-goWeb/admin/model"
	"turan/example-goWeb/admin/mq"
	"turan/example-goWeb/admin/utils"
)




func RegisterEmail(ctx *gin.Context)  {
	email, ok := ctx.GetPostForm("email")
	if !ok {
		ctx.JSON(http.StatusOK,IphoneNotNull)
		return
	}

	code ,ok:= ctx.GetPostForm("code")
	if !ok {
		ctx.JSON(http.StatusOK,CodeNotNull)
		return
	}else {
		//判断过期时间
		isExpire, err := cache.IsExpire(email)
		if err!=nil {
			ctx.JSON(http.StatusOK,err.Error())
			return
		}else {
			if isExpire {
				ctx.JSON(http.StatusOK,CodeExpire)
				return
			}
		}
	}
	passwrod, ok := ctx.GetPostForm("password")
	if !ok {
		ctx.JSON(http.StatusOK,PassWordNotNull)
		return
	}
	//判断验证码正确
	rCode, err := cache.Get(email)
	if err!=nil {
		ctx.JSON(http.StatusOK,err.Error())
		return
	}else {
		if rCode !=code {
			ctx.JSON(http.StatusOK,CodeErr)
			return
		}
	}
	//使用后将验证码过期
	isExpire := cache.SetExpire(email)
	if  !isExpire {
		ctx.JSON(http.StatusOK,isExpire)
		return
	}
	user := &model.User{
		Email:      email,
		Password:   passwrod,
	}
	//插入数据
	err = user.RegisterEmail()
	if err !=nil {
		ctx.JSON(http.StatusOK,err.Error())
		return
	}

	ctx.JSON(http.StatusOK,gin.H{"msg":"注册成功"})
}



// @title 邮箱登录
// @Summary 邮箱登录
// @Tags 用户
// @Accept application/json
// @Produce  json
// @Param email body string true "邮箱"
// @Param code body string true "邮箱"
// @Success 200 {object} string "成功"
// @Router /emailLogin [post]
func emailLogin(ctx *gin.Context)  {
	//参数校验
	email, ok := ctx.GetPostForm("email")
	if !ok {
		ctx.JSON(http.StatusOK,IphoneNotNull)
		return
	}
	password ,ok:= ctx.GetPostForm("password")
	if !ok {
		ctx.JSON(http.StatusOK,PassWordNotNull)
		return
	}
	var user = &model.User{}
	user.Email = email
	err := user.EmailLogin()
	if err != nil {
		ctx.JSON(http.StatusOK,UserNotExsits)
		return
	}

	if user.Password != password{
		ctx.JSON(http.StatusOK,PasswordErr)
		return
	}

	//生成token
	token, err := utils.CreateToken(email)
	if err != nil {
		ctx.JSON(http.StatusOK,err.Error())
		return
	}
	//存储otken
	keyName :=fmt.Sprintf("%S%S","token",email)
	err = cache.SetAndTime(keyName, token, utils.TokenExpiresTime)
	if err != nil {
		ctx.JSON(http.StatusOK,err.Error())
		return
	}

	ctx.JSON(http.StatusOK,gin.H{"msg":"登录成功","token":token})



}



// @title 发送手机验证码
// @Summary 发送手机验证码
// @Tags 用户
// @Accept application/json
// @Produce  json
// @Param email body string true "邮箱"
// @Success 200 {object} string "成功"
// @Failure 200 {object} Code "请求错误"
// @Failure 500 {object} Code "内部错误"
// @Router /sendEmailCode [post]
func sendEmailCode(ctx *gin.Context)  {
	//参数校验
	email,ok := ctx.GetPostForm("email")
	if !ok {
		ctx.JSON(http.StatusOK,EmaillNotNull)
		return
	}

	//生产验证码
	code := utils.GetCode()

	emailAndCode := utils.Email{email,code}
	//发送断行

	go  mq.EmailQueue(emailAndCode)



	//存储验证码
	err := cache.SetAndTime(email, code,cache.TENMIN)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,err.Error())
		return
	}

	//返回参数
	ctx.JSON(http.StatusOK,gin.H{"msg":"邮箱验证码发送成功！"})

}
