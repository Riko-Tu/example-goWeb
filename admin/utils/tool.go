package utils

import (
	"fmt"
	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"
	"math/rand"
	"strconv"
	"time"
)

//获取验证码
func GetCode() string {
	//获取活种
	nano := time.Now().UnixNano()
	//每个种子对应一个随机值
	rnd := rand.New(rand.NewSource(nano))
	return fmt.Sprintf("%06v", rnd.Int31n(1000000))
}

type Email struct {
	Email string `json:"email"`
	Code string  `json:"code"`
}


//发送邮件
func SendEmail(email Email) error {
	from := viper.GetString("smtp.from")
	host := viper.GetString("smtp.qq.host")
	port, _ := strconv.Atoi(viper.GetString("smtp.qq.port"))
	username := viper.GetString("smtp.qq.username")
	password := viper.GetString("smtp.qq.password")
	m := gomail.NewMessage()                   //获取邮件对象
	m.SetHeader("From", "<"+from+">") //发件人邮箱
	m.SetHeader("To", email.Email)            //收件人邮箱
	m.SetHeader("Subject", "绽放【邮箱验证码】")        //标题
	m.SetBody("text/html", fmt.Sprintf("你邮箱登录的验证码是%s，有效时间十分钟，使用后过期", email.Code))
	//创建smtp拨号器
	d := gomail.Dialer{Host: host, Port: port, Username: username, Password: password}
	//使用拨号器发送message
	return d.DialAndSend(m)
}


