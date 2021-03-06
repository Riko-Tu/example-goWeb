package utils

// This file is auto-generated, don't edit it. Thanks.

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v2/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/spf13/viper"
)

/**
 * 使用AK&SK初始化账号Client
 * @param accessKeyId
 * @param accessKeySecret
 * @return Client
 * @throws Exception
 */
func CreateClient(accessKeyId *string, accessKeySecret *string) (_result *dysmsapi20170525.Client, _err error) {

	config := &openapi.Config{
		// 您的 AccessKey ID
		AccessKeyId: accessKeyId,
		// 您的 AccessKey Secret
		AccessKeySecret: accessKeySecret,
	}
	// 访问的域名
	config.Endpoint = tea.String("dysmsapi.aliyuncs.com")
	_result = &dysmsapi20170525.Client{}
	_result, _err = dysmsapi20170525.NewClient(config)
	return _result, _err
}

func SendCode(args []*string) (_err error) {
	key :=viper.GetString("aliyun.key")
	id := viper.GetString("aliyun.id")

	client, _err := CreateClient(tea.String(id), tea.String(key))
	if _err != nil {
		return _err
	}
	iphone := "15889978676"
	sendSmsRequest := &dysmsapi20170525.SendSmsRequest{
		SignName:      tea.String("zhan"),
		TemplateCode:  tea.String("SMS_154950909"),
		PhoneNumbers:  tea.String(iphone),
		TemplateParam: tea.String("{\"code\":\"222\"}"),
	}
	runtime := &util.RuntimeOptions{}
	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		// 复制代码运行请自行打印 API 的返回值
		_, _err = client.SendSmsWithOptions(sendSmsRequest, runtime)
		if _err != nil {
			return _err
		}

		return nil
	}()

	if tryErr != nil {
		var error = &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			error = _t
		} else {
			error.Message = tea.String(tryErr.Error())
		}
		// 如有需要，请打印 error
		util.AssertAsString(error.Message)
	}
	return _err
}
//err := utils.SendCode(tea.StringSlice(os.Args[1:]))
//if err != nil {
//panic(err)
//}
