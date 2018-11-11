package http_wrappers

import "net/http"

//go:generate moq -out httpClient_mock.go . HttpClient
type HttpClient interface {
	Get(url string) (*http.Response, error)
}
