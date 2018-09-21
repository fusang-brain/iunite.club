package utils

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// SendMsg 发送短信（使用Badsms接口）
func SendMsg(mobile string, content string) bool {

	api := "http://sms.253.com/msg/send"
	user := "N526104_N1238321"
	pwd := "oDZe1OByGH5d5f"
	url := fmt.Sprintf("%s?un=%s&pw=%s&phone=%s&msg=%s&rd=1", api, user, pwd, mobile, content)
	fmt.Println("请求地址:" + url)
	return httpGet(url) != nil
}

func httpGet(url string) []byte {
	r, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("http.NewRequest:", err.Error())
		return nil
	}

	fmt.Println(r.Proto)

	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		fmt.Println("http.DefaultClient.Do:", err.Error())
		return nil
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Println("resp.StatusCode not ok", resp.StatusCode)
		return nil
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil && err != io.EOF {
		fmt.Println("ioutil.ReadAll:", err.Error())
		return nil
	}

	fmt.Println(string(data))
	return data
}
