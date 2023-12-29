package reqclient

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func HttpRequest(method, rawUrl string, bodyMaps, headers map[string]string, timeout time.Duration, jsons string) (result string, err error) {
	var (
		request  *http.Request
		response *http.Response
		res      []byte
	)
	if timeout <= 0 {
		timeout = 5
	}
	client := &http.Client{
		Timeout: timeout * time.Second,
	}

	if jsons == "" {
		// 请求的 body 内容
		data := url.Values{}
		for key, value := range bodyMaps {
			data.Set(key, value)
		}

		jsons = data.Encode()
	}

	if request, err = http.NewRequest(method, rawUrl, strings.NewReader(jsons)); err != nil {
		return
	}

	if method == "GET" {
		request.URL.RawQuery = jsons
	}

	// 提交请求
	// 增加header头信息
	for key, val := range headers {
		request.Header.Set(key, val)
	}
	// 处理返回结果
	if response, err = client.Do(request); err != nil {
		return "", fmt.Errorf("clinet do 关闭了 %s", err)
	}

	//这里判断下，如果响应关闭。则直接返回，实测是存在这种情况的
	if response == nil || response.Close {
		return "", fmt.Errorf("client has been closed")
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("get content failed status code is %d ", response.StatusCode)
	}
	if res, err = io.ReadAll(response.Body); err != nil {
		return
	}
	return string(res), nil
}
