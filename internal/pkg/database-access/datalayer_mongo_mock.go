// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package database_access

import (
	"sync"
)

var (
	lockDataLayerMockC sync.RWMutex
)

// DataLayerMock is a mock implementation of DataLayer.
//
//     func TestSomethingThatUsesDataLayer(t *testing.T) {
//
//         // make and configure a mocked DataLayer
//         mockedDataLayer := &DataLayerMock{
//             CFunc: func(name string) Collection {
// 	               panic("mock out the C method")
//             },
//         }
//
//         // use mockedDataLayer in code that requires DataLayer
//         // and then make assertions.
//
//     }
type DataLayerMock struct {
	// CFunc mocks the C method.
	CFunc func(name string) Collection

	// calls tracks calls to the methods.
	calls struct {
		// C holds details about calls to the C method.
		C []struct {
			// Name is the name argument value.
			Name string
		}
	}
}

// C calls CFunc.
func (mock *DataLayerMock) C(name string) Collection {
	if mock.CFunc == nil {
		panic("DataLayerMock.CFunc: method is nil but DataLayer.C was just called")
	}
	callInfo := struct {
		Name string
	}{
		Name: name,
	}
	lockDataLayerMockC.Lock()
	mock.calls.C = append(mock.calls.C, callInfo)
	lockDataLayerMockC.Unlock()
	return mock.CFunc(name)
}

// CCalls gets all the calls that were made to C.
// Check the length with:
//     len(mockedDataLayer.CCalls())
func (mock *DataLayerMock) CCalls() []struct {
	Name string
} {
	var calls []struct {
		Name string
	}
	lockDataLayerMockC.RLock()
	calls = mock.calls.C
	lockDataLayerMockC.RUnlock()
	return calls
}
