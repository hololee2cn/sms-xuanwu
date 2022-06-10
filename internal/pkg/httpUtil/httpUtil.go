// +build go1.13

package httpUtil

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// type client struct {
// 	client *http.Client
// }

// var (
// clients          = make(map[string]*client)
// errInvaildMethod = errors.New("invalid request method")
// errInvaildUrl    = errors.New("request url is empty")
// )

// func init() {
// 	// 初始化一个client,重复使用
// 	//
// 	normalClient := &client{
// 		client: &http.Client{
// 			Timeout: time.Second * 30,
// 		},
// 	}
// 	clients["normal"] = normalClient
// }

// func (c *client) Do(req *http.Request) (*http.Response, error) {
// 	return c.client.Do(req)
// }

func Request(method, url string, body []byte, header map[string][]string) (int, []byte, error) {
	return request(http.DefaultClient, method, url, body, header)
}

func request(client *http.Client, method, u string, body []byte, header map[string][]string) (
	code int, res []byte, err error) {
	// 判断请求的方法是否正确
	switch method {
	case http.MethodGet, http.MethodHead, http.MethodPost, http.MethodPut, http.MethodPatch,
		http.MethodDelete, http.MethodConnect, http.MethodOptions, http.MethodTrace:
	default:
		err = fmt.Errorf("go [beep] yourself if you're using customer http method: %s\n", method)
		return
	}

	// 判断url是否有效
	if _, err = url.ParseRequestURI(u); err != nil {
		err = fmt.Errorf("invalid request url, error: %w url: %v", err, u)
		return
	}

	// 创建一个request
	r, err := http.NewRequest(method, u, bytes.NewReader(body))
	if err != nil {
		err = fmt.Errorf("make request fail, error: %w, method: %v, url: %v, body: %s", err, method, u, body)
		return
	}

	// 观察一下，看看这么用是不是好一点点（主要是为了转发header）
	if len(header) > 0 {
		r.Header = header
	}
	// if len(header) > 0 {
	// 	for k, v := range header {
	// 		if "" != k {
	// 			r.Header.Set(k, v)
	// 		}
	// 	}
	// }

	// 设置代理信息
	r.Header.Set("User-Agent", "test")

	resp, err := client.Do(r)
	if err != nil {
		err = fmt.Errorf("failed to send request, err: %w", err)
		return
	}

	// 限制响应消息体,防止获得恶意攻击信息
	// resp.BodybodyRead := http.MaxBytesReader(nil, resp.Body, 1024*1024*10)

	code = resp.StatusCode
	// defer bodyRead.Close()
	defer resp.Body.Close()
	res, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		err = fmt.Errorf("failed to read response body, err: %w", err)
		return
	}

	return
}

// NewRequest 新建http请求
func NewRequest(method, url, body string, header map[string]string) (*http.Request, error) {
	request, err := http.NewRequest(method, url, bytes.NewBufferString(body))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	request.Header.Set("User-Agent", "test")

	for key, value := range header {
		if key != "" {
			request.Header.Set(key, value)
		}
	}

	return request, nil
}

// Fetch 访问API并解析返回的json响应
func Fetch(request *http.Request, bean interface{}) error {
	// 设置代理信息
	request.Header.Set("User-Agent", "test")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		err = fmt.Errorf("failed to get response, error: %w", err)
		fmt.Println(err)
		return err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		err = fmt.Errorf("failed to get response, code: %d, status: %s ", response.StatusCode, response.Status)
		fmt.Println(err)
		return err
	}

	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		err = fmt.Errorf("failed to read response, error: %w", err)
		fmt.Println(err)
		return err
	}

	err = json.Unmarshal(b, &bean)
	if err != nil {
		err = fmt.Errorf("failed to unmarshal response, error: %w", err)
		fmt.Println(err)
		return err
	}

	return nil
}

// 统一http返回格式
type httpResp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func NewResp(code int, msg string, data interface{}) httpResp {
	return httpResp{Code: code, Msg: msg, Data: data}
}

func NewFailedResp(msg string, code ...int) httpResp {
	c := 1
	if len(code) > 0 {
		c = code[0]
	}
	return NewResp(c, msg, nil)
}

// 成功code为0, msg和data都为可选
func NewSuccessResp(msg string, data interface{}) httpResp {
	return NewResp(0, msg, data)
}
