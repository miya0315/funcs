package reqclient

import (
	"encoding/json"
	"log"
	"testing"
)

func TestRequest(t *testing.T) {
	url := "http://localhost:8208/api/v1/push/getway?q=21321&q2=wewqe21"
	bd := map[string]string{"first": "123123213122"}
	jss,_ := json.Marshal(bd)
	js, err := HttpRequest("POST", url, nil, map[string]string{"Content-Type": "application/x-www-form-urlencoded"}, 5, string(jss))
	log.Println(js, err)
}
