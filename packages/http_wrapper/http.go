package http_wrapper

import (
	"io"
	"net/http"
)

type IHttp interface {
	Do(req *http.Request) (*http.Response, error)
	Get(url string) (resp *http.Response, err error)
	NewRequest(method, url string, body io.Reader) (*http.Request, error)
}

var HttpService = httpServiceImpl{}

// implementation
type httpServiceImpl struct{}

func (httpServiceImpl) Do(req *http.Request) (*http.Response, error) {
	return http.DefaultClient.Do(req)
}
func (httpServiceImpl) Get(url string) (resp *http.Response, err error) {
	return http.Get(url)
}

func (httpServiceImpl) NewRequest(method, url string, body io.Reader) (*http.Request, error) {
	return http.NewRequest(method, url, body)
}
