package http_wrapper_test

import (
	"io"
	"net/http"
)

// mock implementation
type MockHttpClient struct {
	DoFuncMock         func() (*http.Response, error)
	GetFuncMock        func() (resp *http.Response, err error)
	NewRequestMockFunc func(method, url string, body io.Reader) (*http.Request, error)
}

func (m MockHttpClient) Do(req *http.Request) (*http.Response, error) {
	return m.DoFuncMock()
}
func (m MockHttpClient) Get(url string) (resp *http.Response, err error) {
	return m.GetFuncMock()
}

func (m MockHttpClient) NewRequest(method, url string, body io.Reader) (*http.Request, error) {
	return m.NewRequestMockFunc(method, url, body)
}
