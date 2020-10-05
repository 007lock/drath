// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package contract

import (
	"context"
	"sync"
)

var (
	lockCursorPaginatorMockGetAfterCursor  sync.RWMutex
	lockCursorPaginatorMockGetBeforeCursor sync.RWMutex
	lockCursorPaginatorMockNums            sync.RWMutex
	lockCursorPaginatorMockResults         sync.RWMutex
	lockCursorPaginatorMockSetAfterCursor  sync.RWMutex
	lockCursorPaginatorMockSetBeforeCursor sync.RWMutex
	lockCursorPaginatorMockSetLimit        sync.RWMutex
	lockCursorPaginatorMockSetOrderBy      sync.RWMutex
)

// Ensure, that CursorPaginatorMock does implement CursorPaginator.
// If this is not the case, regenerate this file with moq.
var _ CursorPaginator = &CursorPaginatorMock{}

// CursorPaginatorMock is a mock implementation of CursorPaginator.
//
//     func TestSomethingThatUsesCursorPaginator(t *testing.T) {
//
//         // make and configure a mocked CursorPaginator
//         mockedCursorPaginator := &CursorPaginatorMock{
//             GetAfterCursorFunc: func() string {
// 	               panic("mock out the GetAfterCursor method")
//             },
//             GetBeforeCursorFunc: func() string {
// 	               panic("mock out the GetBeforeCursor method")
//             },
//             NumsFunc: func() uint64 {
// 	               panic("mock out the Nums method")
//             },
//             ResultsFunc: func(ctx context.Context, table string, data interface{}) error {
// 	               panic("mock out the Results method")
//             },
//             SetAfterCursorFunc: func(cursor string)  {
// 	               panic("mock out the SetAfterCursor method")
//             },
//             SetBeforeCursorFunc: func(cursor string)  {
// 	               panic("mock out the SetBeforeCursor method")
//             },
//             SetLimitFunc: func(limit uint64)  {
// 	               panic("mock out the SetLimit method")
//             },
//             SetOrderByFunc: func(key string, order string)  {
// 	               panic("mock out the SetOrderBy method")
//             },
//         }
//
//         // use mockedCursorPaginator in code that requires CursorPaginator
//         // and then make assertions.
//
//     }
type CursorPaginatorMock struct {
	// GetAfterCursorFunc mocks the GetAfterCursor method.
	GetAfterCursorFunc func() string

	// GetBeforeCursorFunc mocks the GetBeforeCursor method.
	GetBeforeCursorFunc func() string

	// NumsFunc mocks the Nums method.
	NumsFunc func() uint64

	// ResultsFunc mocks the Results method.
	ResultsFunc func(ctx context.Context, table string, data interface{}) error

	// SetAfterCursorFunc mocks the SetAfterCursor method.
	SetAfterCursorFunc func(cursor string)

	// SetBeforeCursorFunc mocks the SetBeforeCursor method.
	SetBeforeCursorFunc func(cursor string)

	// SetLimitFunc mocks the SetLimit method.
	SetLimitFunc func(limit uint64)

	// SetOrderByFunc mocks the SetOrderBy method.
	SetOrderByFunc func(key string, order string)

	// calls tracks calls to the methods.
	calls struct {
		// GetAfterCursor holds details about calls to the GetAfterCursor method.
		GetAfterCursor []struct {
		}
		// GetBeforeCursor holds details about calls to the GetBeforeCursor method.
		GetBeforeCursor []struct {
		}
		// Nums holds details about calls to the Nums method.
		Nums []struct {
		}
		// Results holds details about calls to the Results method.
		Results []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Table is the table argument value.
			Table string
			// Data is the data argument value.
			Data interface{}
		}
		// SetAfterCursor holds details about calls to the SetAfterCursor method.
		SetAfterCursor []struct {
			// Cursor is the cursor argument value.
			Cursor string
		}
		// SetBeforeCursor holds details about calls to the SetBeforeCursor method.
		SetBeforeCursor []struct {
			// Cursor is the cursor argument value.
			Cursor string
		}
		// SetLimit holds details about calls to the SetLimit method.
		SetLimit []struct {
			// Limit is the limit argument value.
			Limit uint64
		}
		// SetOrderBy holds details about calls to the SetOrderBy method.
		SetOrderBy []struct {
			// Key is the key argument value.
			Key string
			// Order is the order argument value.
			Order string
		}
	}
}

// GetAfterCursor calls GetAfterCursorFunc.
func (mock *CursorPaginatorMock) GetAfterCursor() string {
	if mock.GetAfterCursorFunc == nil {
		panic("CursorPaginatorMock.GetAfterCursorFunc: method is nil but CursorPaginator.GetAfterCursor was just called")
	}
	callInfo := struct {
	}{}
	lockCursorPaginatorMockGetAfterCursor.Lock()
	mock.calls.GetAfterCursor = append(mock.calls.GetAfterCursor, callInfo)
	lockCursorPaginatorMockGetAfterCursor.Unlock()
	return mock.GetAfterCursorFunc()
}

// GetAfterCursorCalls gets all the calls that were made to GetAfterCursor.
// Check the length with:
//     len(mockedCursorPaginator.GetAfterCursorCalls())
func (mock *CursorPaginatorMock) GetAfterCursorCalls() []struct {
} {
	var calls []struct {
	}
	lockCursorPaginatorMockGetAfterCursor.RLock()
	calls = mock.calls.GetAfterCursor
	lockCursorPaginatorMockGetAfterCursor.RUnlock()
	return calls
}

// GetBeforeCursor calls GetBeforeCursorFunc.
func (mock *CursorPaginatorMock) GetBeforeCursor() string {
	if mock.GetBeforeCursorFunc == nil {
		panic("CursorPaginatorMock.GetBeforeCursorFunc: method is nil but CursorPaginator.GetBeforeCursor was just called")
	}
	callInfo := struct {
	}{}
	lockCursorPaginatorMockGetBeforeCursor.Lock()
	mock.calls.GetBeforeCursor = append(mock.calls.GetBeforeCursor, callInfo)
	lockCursorPaginatorMockGetBeforeCursor.Unlock()
	return mock.GetBeforeCursorFunc()
}

// GetBeforeCursorCalls gets all the calls that were made to GetBeforeCursor.
// Check the length with:
//     len(mockedCursorPaginator.GetBeforeCursorCalls())
func (mock *CursorPaginatorMock) GetBeforeCursorCalls() []struct {
} {
	var calls []struct {
	}
	lockCursorPaginatorMockGetBeforeCursor.RLock()
	calls = mock.calls.GetBeforeCursor
	lockCursorPaginatorMockGetBeforeCursor.RUnlock()
	return calls
}

// Nums calls NumsFunc.
func (mock *CursorPaginatorMock) Nums() uint64 {
	if mock.NumsFunc == nil {
		panic("CursorPaginatorMock.NumsFunc: method is nil but CursorPaginator.Nums was just called")
	}
	callInfo := struct {
	}{}
	lockCursorPaginatorMockNums.Lock()
	mock.calls.Nums = append(mock.calls.Nums, callInfo)
	lockCursorPaginatorMockNums.Unlock()
	return mock.NumsFunc()
}

// NumsCalls gets all the calls that were made to Nums.
// Check the length with:
//     len(mockedCursorPaginator.NumsCalls())
func (mock *CursorPaginatorMock) NumsCalls() []struct {
} {
	var calls []struct {
	}
	lockCursorPaginatorMockNums.RLock()
	calls = mock.calls.Nums
	lockCursorPaginatorMockNums.RUnlock()
	return calls
}

// Results calls ResultsFunc.
func (mock *CursorPaginatorMock) Results(ctx context.Context, table string, data interface{}) error {
	if mock.ResultsFunc == nil {
		panic("CursorPaginatorMock.ResultsFunc: method is nil but CursorPaginator.Results was just called")
	}
	callInfo := struct {
		Ctx   context.Context
		Table string
		Data  interface{}
	}{
		Ctx:   ctx,
		Table: table,
		Data:  data,
	}
	lockCursorPaginatorMockResults.Lock()
	mock.calls.Results = append(mock.calls.Results, callInfo)
	lockCursorPaginatorMockResults.Unlock()
	return mock.ResultsFunc(ctx, table, data)
}

// ResultsCalls gets all the calls that were made to Results.
// Check the length with:
//     len(mockedCursorPaginator.ResultsCalls())
func (mock *CursorPaginatorMock) ResultsCalls() []struct {
	Ctx   context.Context
	Table string
	Data  interface{}
} {
	var calls []struct {
		Ctx   context.Context
		Table string
		Data  interface{}
	}
	lockCursorPaginatorMockResults.RLock()
	calls = mock.calls.Results
	lockCursorPaginatorMockResults.RUnlock()
	return calls
}

// SetAfterCursor calls SetAfterCursorFunc.
func (mock *CursorPaginatorMock) SetAfterCursor(cursor string) {
	if mock.SetAfterCursorFunc == nil {
		panic("CursorPaginatorMock.SetAfterCursorFunc: method is nil but CursorPaginator.SetAfterCursor was just called")
	}
	callInfo := struct {
		Cursor string
	}{
		Cursor: cursor,
	}
	lockCursorPaginatorMockSetAfterCursor.Lock()
	mock.calls.SetAfterCursor = append(mock.calls.SetAfterCursor, callInfo)
	lockCursorPaginatorMockSetAfterCursor.Unlock()
	mock.SetAfterCursorFunc(cursor)
}

// SetAfterCursorCalls gets all the calls that were made to SetAfterCursor.
// Check the length with:
//     len(mockedCursorPaginator.SetAfterCursorCalls())
func (mock *CursorPaginatorMock) SetAfterCursorCalls() []struct {
	Cursor string
} {
	var calls []struct {
		Cursor string
	}
	lockCursorPaginatorMockSetAfterCursor.RLock()
	calls = mock.calls.SetAfterCursor
	lockCursorPaginatorMockSetAfterCursor.RUnlock()
	return calls
}

// SetBeforeCursor calls SetBeforeCursorFunc.
func (mock *CursorPaginatorMock) SetBeforeCursor(cursor string) {
	if mock.SetBeforeCursorFunc == nil {
		panic("CursorPaginatorMock.SetBeforeCursorFunc: method is nil but CursorPaginator.SetBeforeCursor was just called")
	}
	callInfo := struct {
		Cursor string
	}{
		Cursor: cursor,
	}
	lockCursorPaginatorMockSetBeforeCursor.Lock()
	mock.calls.SetBeforeCursor = append(mock.calls.SetBeforeCursor, callInfo)
	lockCursorPaginatorMockSetBeforeCursor.Unlock()
	mock.SetBeforeCursorFunc(cursor)
}

// SetBeforeCursorCalls gets all the calls that were made to SetBeforeCursor.
// Check the length with:
//     len(mockedCursorPaginator.SetBeforeCursorCalls())
func (mock *CursorPaginatorMock) SetBeforeCursorCalls() []struct {
	Cursor string
} {
	var calls []struct {
		Cursor string
	}
	lockCursorPaginatorMockSetBeforeCursor.RLock()
	calls = mock.calls.SetBeforeCursor
	lockCursorPaginatorMockSetBeforeCursor.RUnlock()
	return calls
}

// SetLimit calls SetLimitFunc.
func (mock *CursorPaginatorMock) SetLimit(limit uint64) {
	if mock.SetLimitFunc == nil {
		panic("CursorPaginatorMock.SetLimitFunc: method is nil but CursorPaginator.SetLimit was just called")
	}
	callInfo := struct {
		Limit uint64
	}{
		Limit: limit,
	}
	lockCursorPaginatorMockSetLimit.Lock()
	mock.calls.SetLimit = append(mock.calls.SetLimit, callInfo)
	lockCursorPaginatorMockSetLimit.Unlock()
	mock.SetLimitFunc(limit)
}

// SetLimitCalls gets all the calls that were made to SetLimit.
// Check the length with:
//     len(mockedCursorPaginator.SetLimitCalls())
func (mock *CursorPaginatorMock) SetLimitCalls() []struct {
	Limit uint64
} {
	var calls []struct {
		Limit uint64
	}
	lockCursorPaginatorMockSetLimit.RLock()
	calls = mock.calls.SetLimit
	lockCursorPaginatorMockSetLimit.RUnlock()
	return calls
}

// SetOrderBy calls SetOrderByFunc.
func (mock *CursorPaginatorMock) SetOrderBy(key string, order string) {
	if mock.SetOrderByFunc == nil {
		panic("CursorPaginatorMock.SetOrderByFunc: method is nil but CursorPaginator.SetOrderBy was just called")
	}
	callInfo := struct {
		Key   string
		Order string
	}{
		Key:   key,
		Order: order,
	}
	lockCursorPaginatorMockSetOrderBy.Lock()
	mock.calls.SetOrderBy = append(mock.calls.SetOrderBy, callInfo)
	lockCursorPaginatorMockSetOrderBy.Unlock()
	mock.SetOrderByFunc(key, order)
}

// SetOrderByCalls gets all the calls that were made to SetOrderBy.
// Check the length with:
//     len(mockedCursorPaginator.SetOrderByCalls())
func (mock *CursorPaginatorMock) SetOrderByCalls() []struct {
	Key   string
	Order string
} {
	var calls []struct {
		Key   string
		Order string
	}
	lockCursorPaginatorMockSetOrderBy.RLock()
	calls = mock.calls.SetOrderBy
	lockCursorPaginatorMockSetOrderBy.RUnlock()
	return calls
}