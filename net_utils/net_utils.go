package net_utils

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

// HTTP GET 请求
func Get(url string) string {

	res, _ := http.Get(url)
	reader, _ := ioutil.ReadAll(res.Body)
	return string(reader)
}
func A(url string) *http.Response {
	//生成client 参数为默认
	client := &http.Client{}
	//提交请求
	reqest, _ := http.NewRequest("GET", url, nil)
	//处理返回结果
	response, _ := client.Do(reqest)
	//将结果定位到标准输出 也可以直接打印出来 或者定位到其他地方进行相应的处理
	stdout := os.Stdout
	//输出请求体
	_, _ = io.Copy(stdout, response.Body)

	return response
}
