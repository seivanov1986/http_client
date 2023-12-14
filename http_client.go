package http_client

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

type httpClient struct {
	Url       string
	Method    string
	PostData  io.Reader
	Header    map[string][]string
	Response  []byte
	Cookies   []*http.Cookie
	BasicAuth string
	Status    string
}

func New() *httpClient {
	return &httpClient{}
}

func (c *httpClient) GetStatus() string {
	return c.Status
}

func (c *httpClient) GetResponse() []byte {
	return c.Response
}

func (c *httpClient) SetUrl(url string) {
	c.Url = url
}

func (c *httpClient) SetMethod(method string) {
	c.Method = method
}

func (c *httpClient) SetPostByteData(postData []byte) {
	c.PostData = bytes.NewReader(postData)
}

func (c *httpClient) SetPostData(postData string) {
	c.PostData = strings.NewReader(postData)
}

func (c *httpClient) SetHeader(header map[string][]string) {
	c.Header = header
}

func (c *httpClient) SetCookies(cookies []*http.Cookie) {
	c.Cookies = cookies
}

func (c *httpClient) SetBasicAuth(username string, password string) {
	auth := username + ":" + password
	c.BasicAuth = base64.StdEncoding.EncodeToString([]byte(auth))
}

func (c *httpClient) Exec() {
	c.Response = nil

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}
	req, _ := http.NewRequest(c.Method, c.Url, c.PostData)
	req.Header = c.Header

	if len(c.BasicAuth) > 0 {
		req.Header.Add("Authorization", "Basic "+c.BasicAuth)
	}

	for _, v := range c.Cookies {
		req.AddCookie(v)
	}

	resp, err := client.Do(req)
	c.Status = resp.Status

	var data []byte
	if err == nil {
		data, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			c.Response = nil
			return
		}
	} else {
		fmt.Println(err)
		c.Response = nil
		return
	}

	c.Response = data
}
