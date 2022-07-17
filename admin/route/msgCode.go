package route



type Code struct {
	StatusCode  int
	Err string
}



//一级参数错误
var (
	IphoneNotNull = Code{10001, "电话号码不能为空"}

	EmaillNotNull = Code{10002,"邮箱不能为空"}

	CodeNotNull = Code{10003,"验证码不能为空"}

	CodeErr =Code{10004,"验证码错误"}

	CodeExpire =Code{10005,"验证码已过期"}

	PassWordNotNull = Code{10006,"密码不能为空"}

	PasswordErr  =Code{10007,"密码错误"}

	UserNotExsits = Code{10008,"用户不存在"}

)

//二非具体
var(
	ServerBusy = Code{20000,"服务器繁忙,请稍后再试"}
)
