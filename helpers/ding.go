package helpers

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func AlertToDingDing(message string) error {
	data := map[string]interface{}{
		"msgtype": "text",
		"text":    map[string]string{"content": message},
	}
	postBody, _ := json.Marshal(data)
	buffer := bytes.NewBuffer(postBody)
	request, err := http.NewRequest("POST", "https://oapi.dingtalk.com/robot/send?access_token=59dfc1b5e907176c5a07eca732cc986e6c467a7a504eb71ba009a83542c59098", buffer)
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8") //添加请求头
	client := http.Client{}
	//创建客户端
	resp, err := client.Do(request.WithContext(context.TODO())) //发送请求
	if err != nil {
		return err
	}
	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return nil
}
