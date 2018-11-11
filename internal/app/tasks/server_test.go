package tasks

import (
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time-logger/internal/pkg/http-wrappers"
	"time-logger/internal/pkg/server"
)

func TestStartServer(t *testing.T) {
	mockedRouter := &server.RouterMock{
		HandleFunc: func(path string, handler http.Handler) *mux.Route {
			return &mux.Route{}
		},
		MethodsFunc: func(methods ...string) *mux.Route {
			return &mux.Route{}
		},
		ServeHTTPFunc: func(w http.ResponseWriter, r *http.Request)  {
			return
		},
	}

	StartServer(mockedRouter, &http_wrappers.Env{})

	hc := mockedRouter.HandleCalls()
	assert.Equal(t, 5, len(hc))

}