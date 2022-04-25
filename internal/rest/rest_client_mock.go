package rest

import "net/http"

var GetFunc func(url string) (*http.Response, error)

type MockClient struct {
	GetFunc func(url string) (*http.Response, error)
}

func (m *MockClient) Get(url string) (*http.Response, error) {
	return GetFunc(url)
}
