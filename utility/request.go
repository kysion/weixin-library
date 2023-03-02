package utility

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

// 设置请请求头发起请求
func Request(method, url string, body io.Reader) (string, error) {
	client := &http.Client{
		//设置超时时间
		Timeout: 30 * time.Second,
	}
	// method 请求方法
	// url 路径
	// body 请求体，可创建bytes.NewBuffer()管道
	// 发起一个新的request请求
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		//已经返回了，只要在调用这个函数的时候打印错误即可
		return "", err
	}
	// 设置请求头
	for key, value := range Header() {
		//这里就是循环遍历，设置请求头
		request.Header.Add(key, value)
	}
	//key，value 键值对，所以我用的map装起来，然后遍历，这样看着好看，代码更简洁，
	//这样写太难看了，又多，对的，变量装起来
	//request.Header.Add("Accept","text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	// 客户端发起请求，获取到响应数据
	response, err := client.Do(request)
	if err != nil {
		//fmt.Println(err.Error())
		return "", err
	}
	// 读取获取的响应体
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		//fmt.Println(err.Error())
		return "", err
	}
	// 就是最后在关闭的意思，
	defer response.Body.Close()
	return string(bytes), nil
}

func RequestSend(method string, url string, body io.Reader) (re []byte, err error) {
	//http://apis.juhe.cn/mobile/get?phone=13429667914&key=您申请的KEY
	//创建客户端
	client := http.Client{
		//设置超时时间
		Timeout: 30 * time.Second,
	}
	//创建请求
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("请求创建失败")
		return
	}
	//客户端发起请求
	response, err := client.Do(request)
	if response.StatusCode != 200 {
		fmt.Println("创建客户端请求失败")
		fmt.Println(err.Error())
		return
	}
	//返回的响应进行处理
	fmt.Println("获得响应啦？")
	//读取响应体数据
	bytes, err := ioutil.ReadAll(response.Body)
	fmt.Println(string(bytes))
	//返回响应体数据
	return bytes, nil
}
