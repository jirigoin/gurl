package rest

import "net/http"

var Client HTTPClient

func init() {
	Client = &http.Client{}
}

type HTTPClient interface {
	Get(url string) (*http.Response, error)
}
