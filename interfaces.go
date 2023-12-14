package http_client

import (
	"net/http"
)

type HttpClient interface {
	SetUrl(url string)
	SetMethod(method string)
	SetPostByteData(postData []byte)
	SetPostData(postData string)
	SetHeader(header map[string][]string)
	SetCookies(cookies []*http.Cookie)
	GetResponse() []byte
	GetStatus() string
	SetBasicAuth(username string, password string)
	Exec()
}
