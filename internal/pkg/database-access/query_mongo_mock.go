// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package database_access

import (
	"sync"
)

var (
	lockQueryMockAll      sync.RWMutex
	lockQueryMockDistinct sync.RWMutex
	lockQueryMockOne      sync.RWMutex
)

// QueryMock is a mock implementation of Query.
//
//     func TestSomethingThatUsesQuery(t *testing.T) {
//
//         // make and configure a mocked Query
//         mockedQuery := &QueryMock{
//             AllFunc: func(result interface{}) error {
// 	               panic("mock out the All method")
//             },
//             DistinctFunc: func(key string, result interface{}) error {
// 	               panic("mock out the Distinct method")
//             },
//             OneFunc: func(result interface{}) error {
// 	               panic("mock out the One method")
//             },
//         }
//
//         // use mockedQuery in code that requires Query
//         // and then make assertions.
//
//     }
type QueryMock struct {
	// AllFunc mocks the All method.
	AllFunc func(result interface{}) error

	// DistinctFunc mocks the Distinct method.
	DistinctFunc func(key string, result interface{}) error

	// OneFunc mocks the One method.
	OneFunc func(result interface{}) error

	// calls tracks calls to the methods.
	calls struct {
		// All holds details about calls to the All method.
		All []struct {
			// Result is the result argument value.
			Result interface{}
		}
		// Distinct holds details about calls to the Distinct method.
		Distinct []struct {
			// Key is the key argument value.
			Key string
			// Result is the result argument value.
			Result interface{}
		}
		// One holds details about calls to the One method.
		One []struct {
			// Result is the result argument value.
			Result interface{}
		}
	}
}

// All calls AllFunc.
func (mock *QueryMock) All(result interface{}) error {
	if mock.AllFunc == nil {
		panic("QueryMock.AllFunc: method is nil but Query.All was just called")
	}
	callInfo := struct {
		Result interface{}
	}{
		Result: result,
	}
	lockQueryMockAll.Lock()
	mock.calls.All = append(mock.calls.All, callInfo)
	lockQueryMockAll.Unlock()
	return mock.AllFunc(result)
}

// AllCalls gets all the calls that were made to All.
// Check the length with:
//     len(mockedQuery.AllCalls())
func (mock *QueryMock) AllCalls() []struct {
	Result interface{}
} {
	var calls []struct {
		Result interface{}
	}
	lockQueryMockAll.RLock()
	calls = mock.calls.All
	lockQueryMockAll.RUnlock()
	return calls
}

// Distinct calls DistinctFunc.
func (mock *QueryMock) Distinct(key string, result interface{}) error {
	if mock.DistinctFunc == nil {
		panic("QueryMock.DistinctFunc: method is nil but Query.Distinct was just called")
	}
	callInfo := struct {
		Key    string
		Result interface{}
	}{
		Key:    key,
		Result: result,
	}
	lockQueryMockDistinct.Lock()
	mock.calls.Distinct = append(mock.calls.Distinct, callInfo)
	lockQueryMockDistinct.Unlock()
	return mock.DistinctFunc(key, result)
}

// DistinctCalls gets all the calls that were made to Distinct.
// Check the length with:
//     len(mockedQuery.DistinctCalls())
func (mock *QueryMock) DistinctCalls() []struct {
	Key    string
	Result interface{}
} {
	var calls []struct {
		Key    string
		Result interface{}
	}
	lockQueryMockDistinct.RLock()
	calls = mock.calls.Distinct
	lockQueryMockDistinct.RUnlock()
	return calls
}

// One calls OneFunc.
func (mock *QueryMock) One(result interface{}) error {
	if mock.OneFunc == nil {
		panic("QueryMock.OneFunc: method is nil but Query.One was just called")
	}
	callInfo := struct {
		Result interface{}
	}{
		Result: result,
	}
	lockQueryMockOne.Lock()
	mock.calls.One = append(mock.calls.One, callInfo)
	lockQueryMockOne.Unlock()
	return mock.OneFunc(result)
}

// OneCalls gets all the calls that were made to One.
// Check the length with:
//     len(mockedQuery.OneCalls())
func (mock *QueryMock) OneCalls() []struct {
	Result interface{}
} {
	var calls []struct {
		Result interface{}
	}
	lockQueryMockOne.RLock()
	calls = mock.calls.One
	lockQueryMockOne.RUnlock()
	return calls
}
