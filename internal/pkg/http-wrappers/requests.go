package http_wrappers

import (
	"io/ioutil"
	"net/http"
)

type HTTPResponse struct {
	Status string
	Body   []byte
}

func MakeRequest(url string, ch chan<- HTTPResponse) {
	httpResponse, _ := http.Get(url)
	httpBody, _ := ioutil.ReadAll(httpResponse.Body)
	ch <- HTTPResponse{httpResponse.Status, httpBody}
}
