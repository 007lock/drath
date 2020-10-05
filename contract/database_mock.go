// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package contract

import (
	"sync"
)

var (
	lockDatabaseMockBegin         sync.RWMutex
	lockDatabaseMockClose         sync.RWMutex
	lockDatabaseMockGet           sync.RWMutex
	lockDatabaseMockMigrationDown sync.RWMutex
	lockDatabaseMockMigrationUp   sync.RWMutex
)

// Ensure, that DatabaseMock does implement Database.
// If this is not the case, regenerate this file with moq.
var _ Database = &DatabaseMock{}

// DatabaseMock is a mock implementation of Database.
//
//     func TestSomethingThatUsesDatabase(t *testing.T) {
//
//         // make and configure a mocked Database
//         mockedDatabase := &DatabaseMock{
//             BeginFunc: func() (interface{}, error) {
// 	               panic("mock out the Begin method")
//             },
//             CloseFunc: func() error {
// 	               panic("mock out the Close method")
//             },
//             GetFunc: func() (interface{}, error) {
// 	               panic("mock out the Get method")
//             },
//             MigrationDownFunc: func() error {
// 	               panic("mock out the MigrationDown method")
//             },
//             MigrationUpFunc: func() error {
// 	               panic("mock out the MigrationUp method")
//             },
//         }
//
//         // use mockedDatabase in code that requires Database
//         // and then make assertions.
//
//     }
type DatabaseMock struct {
	// BeginFunc mocks the Begin method.
	BeginFunc func() (interface{}, error)

	// CloseFunc mocks the Close method.
	CloseFunc func() error

	// GetFunc mocks the Get method.
	GetFunc func() (interface{}, error)

	// MigrationDownFunc mocks the MigrationDown method.
	MigrationDownFunc func() error

	// MigrationUpFunc mocks the MigrationUp method.
	MigrationUpFunc func() error

	// calls tracks calls to the methods.
	calls struct {
		// Begin holds details about calls to the Begin method.
		Begin []struct {
		}
		// Close holds details about calls to the Close method.
		Close []struct {
		}
		// Get holds details about calls to the Get method.
		Get []struct {
		}
		// MigrationDown holds details about calls to the MigrationDown method.
		MigrationDown []struct {
		}
		// MigrationUp holds details about calls to the MigrationUp method.
		MigrationUp []struct {
		}
	}
}

// Begin calls BeginFunc.
func (mock *DatabaseMock) Begin() (interface{}, error) {
	if mock.BeginFunc == nil {
		panic("DatabaseMock.BeginFunc: method is nil but Database.Begin was just called")
	}
	callInfo := struct {
	}{}
	lockDatabaseMockBegin.Lock()
	mock.calls.Begin = append(mock.calls.Begin, callInfo)
	lockDatabaseMockBegin.Unlock()
	return mock.BeginFunc()
}

// BeginCalls gets all the calls that were made to Begin.
// Check the length with:
//     len(mockedDatabase.BeginCalls())
func (mock *DatabaseMock) BeginCalls() []struct {
} {
	var calls []struct {
	}
	lockDatabaseMockBegin.RLock()
	calls = mock.calls.Begin
	lockDatabaseMockBegin.RUnlock()
	return calls
}

// Close calls CloseFunc.
func (mock *DatabaseMock) Close() error {
	if mock.CloseFunc == nil {
		panic("DatabaseMock.CloseFunc: method is nil but Database.Close was just called")
	}
	callInfo := struct {
	}{}
	lockDatabaseMockClose.Lock()
	mock.calls.Close = append(mock.calls.Close, callInfo)
	lockDatabaseMockClose.Unlock()
	return mock.CloseFunc()
}

// CloseCalls gets all the calls that were made to Close.
// Check the length with:
//     len(mockedDatabase.CloseCalls())
func (mock *DatabaseMock) CloseCalls() []struct {
} {
	var calls []struct {
	}
	lockDatabaseMockClose.RLock()
	calls = mock.calls.Close
	lockDatabaseMockClose.RUnlock()
	return calls
}

// Get calls GetFunc.
func (mock *DatabaseMock) Get() (interface{}, error) {
	if mock.GetFunc == nil {
		panic("DatabaseMock.GetFunc: method is nil but Database.Get was just called")
	}
	callInfo := struct {
	}{}
	lockDatabaseMockGet.Lock()
	mock.calls.Get = append(mock.calls.Get, callInfo)
	lockDatabaseMockGet.Unlock()
	return mock.GetFunc()
}

// GetCalls gets all the calls that were made to Get.
// Check the length with:
//     len(mockedDatabase.GetCalls())
func (mock *DatabaseMock) GetCalls() []struct {
} {
	var calls []struct {
	}
	lockDatabaseMockGet.RLock()
	calls = mock.calls.Get
	lockDatabaseMockGet.RUnlock()
	return calls
}

// MigrationDown calls MigrationDownFunc.
func (mock *DatabaseMock) MigrationDown() error {
	if mock.MigrationDownFunc == nil {
		panic("DatabaseMock.MigrationDownFunc: method is nil but Database.MigrationDown was just called")
	}
	callInfo := struct {
	}{}
	lockDatabaseMockMigrationDown.Lock()
	mock.calls.MigrationDown = append(mock.calls.MigrationDown, callInfo)
	lockDatabaseMockMigrationDown.Unlock()
	return mock.MigrationDownFunc()
}

// MigrationDownCalls gets all the calls that were made to MigrationDown.
// Check the length with:
//     len(mockedDatabase.MigrationDownCalls())
func (mock *DatabaseMock) MigrationDownCalls() []struct {
} {
	var calls []struct {
	}
	lockDatabaseMockMigrationDown.RLock()
	calls = mock.calls.MigrationDown
	lockDatabaseMockMigrationDown.RUnlock()
	return calls
}

// MigrationUp calls MigrationUpFunc.
func (mock *DatabaseMock) MigrationUp() error {
	if mock.MigrationUpFunc == nil {
		panic("DatabaseMock.MigrationUpFunc: method is nil but Database.MigrationUp was just called")
	}
	callInfo := struct {
	}{}
	lockDatabaseMockMigrationUp.Lock()
	mock.calls.MigrationUp = append(mock.calls.MigrationUp, callInfo)
	lockDatabaseMockMigrationUp.Unlock()
	return mock.MigrationUpFunc()
}

// MigrationUpCalls gets all the calls that were made to MigrationUp.
// Check the length with:
//     len(mockedDatabase.MigrationUpCalls())
func (mock *DatabaseMock) MigrationUpCalls() []struct {
} {
	var calls []struct {
	}
	lockDatabaseMockMigrationUp.RLock()
	calls = mock.calls.MigrationUp
	lockDatabaseMockMigrationUp.RUnlock()
	return calls
}