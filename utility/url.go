package utility

import (
	"fmt"
	"net/url"
)

func URLUtil() {
	//url参数数据格式化  (url编码格式问题)
	param := url.Values{}
	// 设置参数
	//param.Set("consName",userdata.ConsName)
	//param.Set("type",userdata.Type)
	//param.Set("key",conf.HOROKEY)
	var Url *url.URL
	Url, err := url.Parse("https://....")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//如果参数中有中文参数,这个方法会进行URLEncode
	Url.RawQuery = param.Encode()
}
