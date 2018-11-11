// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package http_wrappers

import (
	"net/http"
	"sync"
)

var (
	lockHttpClientMockGet sync.RWMutex
)

// HttpClientMock is a mock implementation of HttpClient.
//
//     func TestSomethingThatUsesHttpClient(t *testing.T) {
//
//         // make and configure a mocked HttpClient
//         mockedHttpClient := &HttpClientMock{
//             GetFunc: func(url string) (*http.Response, error) {
// 	               panic("mock out the Get method")
//             },
//         }
//
//         // use mockedHttpClient in code that requires HttpClient
//         // and then make assertions.
//
//     }
type HttpClientMock struct {
	// GetFunc mocks the Get method.
	GetFunc func(url string) (*http.Response, error)

	// calls tracks calls to the methods.
	calls struct {
		// Get holds details about calls to the Get method.
		Get []struct {
			// URL is the url argument value.
			URL string
		}
	}
}

// Get calls GetFunc.
func (mock *HttpClientMock) Get(url string) (*http.Response, error) {
	if mock.GetFunc == nil {
		panic("HttpClientMock.GetFunc: method is nil but HttpClient.Get was just called")
	}
	callInfo := struct {
		URL string
	}{
		URL: url,
	}
	lockHttpClientMockGet.Lock()
	mock.calls.Get = append(mock.calls.Get, callInfo)
	lockHttpClientMockGet.Unlock()
	return mock.GetFunc(url)
}

// GetCalls gets all the calls that were made to Get.
// Check the length with:
//     len(mockedHttpClient.GetCalls())
func (mock *HttpClientMock) GetCalls() []struct {
	URL string
} {
	var calls []struct {
		URL string
	}
	lockHttpClientMockGet.RLock()
	calls = mock.calls.Get
	lockHttpClientMockGet.RUnlock()
	return calls
}
