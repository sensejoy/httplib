package httplib

import (
	"fmt"
	"testing"
)

func TestCall(t *testing.T) {
	req := Request{
		Method: GET,
		Url:    "https://www.baidu.com/",
	}
	res := Call(req)
	if res.Status != 200 {
		t.Errorf("call fail : %+v, reason:[%s]\n", req, res.Message)
	} else {
		fmt.Printf("call ok: %+v\n", req)
	}
}

func TestMultiCall(t *testing.T) {
	reqs := make(map[interface{}]Request)
	url := "https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=aaa"
	params := `{"touser":"ooo","msgtype":"text","text":{"content":"111"}}`
	reqs["a"] = Request{Method: GET, Url: "https://www.baidu.com", Headers: Hashmap{"user-agent": "go-http"}}
	reqs["b"] = Request{Method: POST, Url: url, Params: params, Type: JSON}
	for idx, res := range MultiCall(reqs, 8) {
		if res.Status < 0 {
			t.Error("call req:", reqs[idx], "fail, reason:", res.Message)
		} else {
			fmt.Println("call req ", reqs[idx], " ok")
		}
	}
}
