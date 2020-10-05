// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package contract

import (
	"context"
	"sync"
)

var (
	lockFetchRepositoryMockApplyCriteria   sync.RWMutex
	lockFetchRepositoryMockFetchByCriteria sync.RWMutex
	lockFetchRepositoryMockFetchByRandom   sync.RWMutex
	lockFetchRepositoryMockFetchCursor     sync.RWMutex
	lockFetchRepositoryMockFetchPagination sync.RWMutex
	lockFetchRepositoryMockGetByCriteria   sync.RWMutex
	lockFetchRepositoryMockGetByID         sync.RWMutex
	lockFetchRepositoryMockGetByRandom     sync.RWMutex
)

// Ensure, that FetchRepositoryMock does implement FetchRepository.
// If this is not the case, regenerate this file with moq.
var _ FetchRepository = &FetchRepositoryMock{}

// FetchRepositoryMock is a mock implementation of FetchRepository.
//
//     func TestSomethingThatUsesFetchRepository(t *testing.T) {
//
//         // make and configure a mocked FetchRepository
//         mockedFetchRepository := &FetchRepositoryMock{
//             ApplyCriteriaFunc: func(c context.Context, crit *RepoCriterias) context.Context {
// 	               panic("mock out the ApplyCriteria method")
//             },
//             FetchByCriteriaFunc: func(c context.Context, table string, item interface{}, crit *RepoCriterias) error {
// 	               panic("mock out the FetchByCriteria method")
//             },
//             FetchByRandomFunc: func(c context.Context, table string, item interface{}, crit *RepoCriterias, limit uint64) error {
// 	               panic("mock out the FetchByRandom method")
//             },
//             FetchCursorFunc: func(c context.Context, table string, item interface{}, p CursorPaginator, crit *RepoCriterias) error {
// 	               panic("mock out the FetchCursor method")
//             },
//             FetchPaginationFunc: func(c context.Context, table string, item interface{}, p OffsetPaginator, crit *RepoCriterias) error {
// 	               panic("mock out the FetchPagination method")
//             },
//             GetByCriteriaFunc: func(c context.Context, table string, item interface{}, crit *RepoCriterias) error {
// 	               panic("mock out the GetByCriteria method")
//             },
//             GetByIDFunc: func(c context.Context, table string, id string, item interface{}, crit *RepoCriterias) error {
// 	               panic("mock out the GetByID method")
//             },
//             GetByRandomFunc: func(c context.Context, table string, item interface{}, crit *RepoCriterias) error {
// 	               panic("mock out the GetByRandom method")
//             },
//         }
//
//         // use mockedFetchRepository in code that requires FetchRepository
//         // and then make assertions.
//
//     }
type FetchRepositoryMock struct {
	// ApplyCriteriaFunc mocks the ApplyCriteria method.
	ApplyCriteriaFunc func(c context.Context, crit *RepoCriterias) context.Context

	// FetchByCriteriaFunc mocks the FetchByCriteria method.
	FetchByCriteriaFunc func(c context.Context, table string, item interface{}, crit *RepoCriterias) error

	// FetchByRandomFunc mocks the FetchByRandom method.
	FetchByRandomFunc func(c context.Context, table string, item interface{}, crit *RepoCriterias, limit uint64) error

	// FetchCursorFunc mocks the FetchCursor method.
	FetchCursorFunc func(c context.Context, table string, item interface{}, p CursorPaginator, crit *RepoCriterias) error

	// FetchPaginationFunc mocks the FetchPagination method.
	FetchPaginationFunc func(c context.Context, table string, item interface{}, p OffsetPaginator, crit *RepoCriterias) error

	// GetByCriteriaFunc mocks the GetByCriteria method.
	GetByCriteriaFunc func(c context.Context, table string, item interface{}, crit *RepoCriterias) error

	// GetByIDFunc mocks the GetByID method.
	GetByIDFunc func(c context.Context, table string, id string, item interface{}, crit *RepoCriterias) error

	// GetByRandomFunc mocks the GetByRandom method.
	GetByRandomFunc func(c context.Context, table string, item interface{}, crit *RepoCriterias) error

	// calls tracks calls to the methods.
	calls struct {
		// ApplyCriteria holds details about calls to the ApplyCriteria method.
		ApplyCriteria []struct {
			// C is the c argument value.
			C context.Context
			// Crit is the crit argument value.
			Crit *RepoCriterias
		}
		// FetchByCriteria holds details about calls to the FetchByCriteria method.
		FetchByCriteria []struct {
			// C is the c argument value.
			C context.Context
			// Table is the table argument value.
			Table string
			// Item is the item argument value.
			Item interface{}
			// Crit is the crit argument value.
			Crit *RepoCriterias
		}
		// FetchByRandom holds details about calls to the FetchByRandom method.
		FetchByRandom []struct {
			// C is the c argument value.
			C context.Context
			// Table is the table argument value.
			Table string
			// Item is the item argument value.
			Item interface{}
			// Crit is the crit argument value.
			Crit *RepoCriterias
			// Limit is the limit argument value.
			Limit uint64
		}
		// FetchCursor holds details about calls to the FetchCursor method.
		FetchCursor []struct {
			// C is the c argument value.
			C context.Context
			// Table is the table argument value.
			Table string
			// Item is the item argument value.
			Item interface{}
			// P is the p argument value.
			P CursorPaginator
			// Crit is the crit argument value.
			Crit *RepoCriterias
		}
		// FetchPagination holds details about calls to the FetchPagination method.
		FetchPagination []struct {
			// C is the c argument value.
			C context.Context
			// Table is the table argument value.
			Table string
			// Item is the item argument value.
			Item interface{}
			// P is the p argument value.
			P OffsetPaginator
			// Crit is the crit argument value.
			Crit *RepoCriterias
		}
		// GetByCriteria holds details about calls to the GetByCriteria method.
		GetByCriteria []struct {
			// C is the c argument value.
			C context.Context
			// Table is the table argument value.
			Table string
			// Item is the item argument value.
			Item interface{}
			// Crit is the crit argument value.
			Crit *RepoCriterias
		}
		// GetByID holds details about calls to the GetByID method.
		GetByID []struct {
			// C is the c argument value.
			C context.Context
			// Table is the table argument value.
			Table string
			// ID is the id argument value.
			ID string
			// Item is the item argument value.
			Item interface{}
			// Crit is the crit argument value.
			Crit *RepoCriterias
		}
		// GetByRandom holds details about calls to the GetByRandom method.
		GetByRandom []struct {
			// C is the c argument value.
			C context.Context
			// Table is the table argument value.
			Table string
			// Item is the item argument value.
			Item interface{}
			// Crit is the crit argument value.
			Crit *RepoCriterias
		}
	}
}

// ApplyCriteria calls ApplyCriteriaFunc.
func (mock *FetchRepositoryMock) ApplyCriteria(c context.Context, crit *RepoCriterias) context.Context {
	if mock.ApplyCriteriaFunc == nil {
		panic("FetchRepositoryMock.ApplyCriteriaFunc: method is nil but FetchRepository.ApplyCriteria was just called")
	}
	callInfo := struct {
		C    context.Context
		Crit *RepoCriterias
	}{
		C:    c,
		Crit: crit,
	}
	lockFetchRepositoryMockApplyCriteria.Lock()
	mock.calls.ApplyCriteria = append(mock.calls.ApplyCriteria, callInfo)
	lockFetchRepositoryMockApplyCriteria.Unlock()
	return mock.ApplyCriteriaFunc(c, crit)
}

// ApplyCriteriaCalls gets all the calls that were made to ApplyCriteria.
// Check the length with:
//     len(mockedFetchRepository.ApplyCriteriaCalls())
func (mock *FetchRepositoryMock) ApplyCriteriaCalls() []struct {
	C    context.Context
	Crit *RepoCriterias
} {
	var calls []struct {
		C    context.Context
		Crit *RepoCriterias
	}
	lockFetchRepositoryMockApplyCriteria.RLock()
	calls = mock.calls.ApplyCriteria
	lockFetchRepositoryMockApplyCriteria.RUnlock()
	return calls
}

// FetchByCriteria calls FetchByCriteriaFunc.
func (mock *FetchRepositoryMock) FetchByCriteria(c context.Context, table string, item interface{}, crit *RepoCriterias) error {
	if mock.FetchByCriteriaFunc == nil {
		panic("FetchRepositoryMock.FetchByCriteriaFunc: method is nil but FetchRepository.FetchByCriteria was just called")
	}
	callInfo := struct {
		C     context.Context
		Table string
		Item  interface{}
		Crit  *RepoCriterias
	}{
		C:     c,
		Table: table,
		Item:  item,
		Crit:  crit,
	}
	lockFetchRepositoryMockFetchByCriteria.Lock()
	mock.calls.FetchByCriteria = append(mock.calls.FetchByCriteria, callInfo)
	lockFetchRepositoryMockFetchByCriteria.Unlock()
	return mock.FetchByCriteriaFunc(c, table, item, crit)
}

// FetchByCriteriaCalls gets all the calls that were made to FetchByCriteria.
// Check the length with:
//     len(mockedFetchRepository.FetchByCriteriaCalls())
func (mock *FetchRepositoryMock) FetchByCriteriaCalls() []struct {
	C     context.Context
	Table string
	Item  interface{}
	Crit  *RepoCriterias
} {
	var calls []struct {
		C     context.Context
		Table string
		Item  interface{}
		Crit  *RepoCriterias
	}
	lockFetchRepositoryMockFetchByCriteria.RLock()
	calls = mock.calls.FetchByCriteria
	lockFetchRepositoryMockFetchByCriteria.RUnlock()
	return calls
}

// FetchByRandom calls FetchByRandomFunc.
func (mock *FetchRepositoryMock) FetchByRandom(c context.Context, table string, item interface{}, crit *RepoCriterias, limit uint64) error {
	if mock.FetchByRandomFunc == nil {
		panic("FetchRepositoryMock.FetchByRandomFunc: method is nil but FetchRepository.FetchByRandom was just called")
	}
	callInfo := struct {
		C     context.Context
		Table string
		Item  interface{}
		Crit  *RepoCriterias
		Limit uint64
	}{
		C:     c,
		Table: table,
		Item:  item,
		Crit:  crit,
		Limit: limit,
	}
	lockFetchRepositoryMockFetchByRandom.Lock()
	mock.calls.FetchByRandom = append(mock.calls.FetchByRandom, callInfo)
	lockFetchRepositoryMockFetchByRandom.Unlock()
	return mock.FetchByRandomFunc(c, table, item, crit, limit)
}

// FetchByRandomCalls gets all the calls that were made to FetchByRandom.
// Check the length with:
//     len(mockedFetchRepository.FetchByRandomCalls())
func (mock *FetchRepositoryMock) FetchByRandomCalls() []struct {
	C     context.Context
	Table string
	Item  interface{}
	Crit  *RepoCriterias
	Limit uint64
} {
	var calls []struct {
		C     context.Context
		Table string
		Item  interface{}
		Crit  *RepoCriterias
		Limit uint64
	}
	lockFetchRepositoryMockFetchByRandom.RLock()
	calls = mock.calls.FetchByRandom
	lockFetchRepositoryMockFetchByRandom.RUnlock()
	return calls
}

// FetchCursor calls FetchCursorFunc.
func (mock *FetchRepositoryMock) FetchCursor(c context.Context, table string, item interface{}, p CursorPaginator, crit *RepoCriterias) error {
	if mock.FetchCursorFunc == nil {
		panic("FetchRepositoryMock.FetchCursorFunc: method is nil but FetchRepository.FetchCursor was just called")
	}
	callInfo := struct {
		C     context.Context
		Table string
		Item  interface{}
		P     CursorPaginator
		Crit  *RepoCriterias
	}{
		C:     c,
		Table: table,
		Item:  item,
		P:     p,
		Crit:  crit,
	}
	lockFetchRepositoryMockFetchCursor.Lock()
	mock.calls.FetchCursor = append(mock.calls.FetchCursor, callInfo)
	lockFetchRepositoryMockFetchCursor.Unlock()
	return mock.FetchCursorFunc(c, table, item, p, crit)
}

// FetchCursorCalls gets all the calls that were made to FetchCursor.
// Check the length with:
//     len(mockedFetchRepository.FetchCursorCalls())
func (mock *FetchRepositoryMock) FetchCursorCalls() []struct {
	C     context.Context
	Table string
	Item  interface{}
	P     CursorPaginator
	Crit  *RepoCriterias
} {
	var calls []struct {
		C     context.Context
		Table string
		Item  interface{}
		P     CursorPaginator
		Crit  *RepoCriterias
	}
	lockFetchRepositoryMockFetchCursor.RLock()
	calls = mock.calls.FetchCursor
	lockFetchRepositoryMockFetchCursor.RUnlock()
	return calls
}

// FetchPagination calls FetchPaginationFunc.
func (mock *FetchRepositoryMock) FetchPagination(c context.Context, table string, item interface{}, p OffsetPaginator, crit *RepoCriterias) error {
	if mock.FetchPaginationFunc == nil {
		panic("FetchRepositoryMock.FetchPaginationFunc: method is nil but FetchRepository.FetchPagination was just called")
	}
	callInfo := struct {
		C     context.Context
		Table string
		Item  interface{}
		P     OffsetPaginator
		Crit  *RepoCriterias
	}{
		C:     c,
		Table: table,
		Item:  item,
		P:     p,
		Crit:  crit,
	}
	lockFetchRepositoryMockFetchPagination.Lock()
	mock.calls.FetchPagination = append(mock.calls.FetchPagination, callInfo)
	lockFetchRepositoryMockFetchPagination.Unlock()
	return mock.FetchPaginationFunc(c, table, item, p, crit)
}

// FetchPaginationCalls gets all the calls that were made to FetchPagination.
// Check the length with:
//     len(mockedFetchRepository.FetchPaginationCalls())
func (mock *FetchRepositoryMock) FetchPaginationCalls() []struct {
	C     context.Context
	Table string
	Item  interface{}
	P     OffsetPaginator
	Crit  *RepoCriterias
} {
	var calls []struct {
		C     context.Context
		Table string
		Item  interface{}
		P     OffsetPaginator
		Crit  *RepoCriterias
	}
	lockFetchRepositoryMockFetchPagination.RLock()
	calls = mock.calls.FetchPagination
	lockFetchRepositoryMockFetchPagination.RUnlock()
	return calls
}

// GetByCriteria calls GetByCriteriaFunc.
func (mock *FetchRepositoryMock) GetByCriteria(c context.Context, table string, item interface{}, crit *RepoCriterias) error {
	if mock.GetByCriteriaFunc == nil {
		panic("FetchRepositoryMock.GetByCriteriaFunc: method is nil but FetchRepository.GetByCriteria was just called")
	}
	callInfo := struct {
		C     context.Context
		Table string
		Item  interface{}
		Crit  *RepoCriterias
	}{
		C:     c,
		Table: table,
		Item:  item,
		Crit:  crit,
	}
	lockFetchRepositoryMockGetByCriteria.Lock()
	mock.calls.GetByCriteria = append(mock.calls.GetByCriteria, callInfo)
	lockFetchRepositoryMockGetByCriteria.Unlock()
	return mock.GetByCriteriaFunc(c, table, item, crit)
}

// GetByCriteriaCalls gets all the calls that were made to GetByCriteria.
// Check the length with:
//     len(mockedFetchRepository.GetByCriteriaCalls())
func (mock *FetchRepositoryMock) GetByCriteriaCalls() []struct {
	C     context.Context
	Table string
	Item  interface{}
	Crit  *RepoCriterias
} {
	var calls []struct {
		C     context.Context
		Table string
		Item  interface{}
		Crit  *RepoCriterias
	}
	lockFetchRepositoryMockGetByCriteria.RLock()
	calls = mock.calls.GetByCriteria
	lockFetchRepositoryMockGetByCriteria.RUnlock()
	return calls
}

// GetByID calls GetByIDFunc.
func (mock *FetchRepositoryMock) GetByID(c context.Context, table string, id string, item interface{}, crit *RepoCriterias) error {
	if mock.GetByIDFunc == nil {
		panic("FetchRepositoryMock.GetByIDFunc: method is nil but FetchRepository.GetByID was just called")
	}
	callInfo := struct {
		C     context.Context
		Table string
		ID    string
		Item  interface{}
		Crit  *RepoCriterias
	}{
		C:     c,
		Table: table,
		ID:    id,
		Item:  item,
		Crit:  crit,
	}
	lockFetchRepositoryMockGetByID.Lock()
	mock.calls.GetByID = append(mock.calls.GetByID, callInfo)
	lockFetchRepositoryMockGetByID.Unlock()
	return mock.GetByIDFunc(c, table, id, item, crit)
}

// GetByIDCalls gets all the calls that were made to GetByID.
// Check the length with:
//     len(mockedFetchRepository.GetByIDCalls())
func (mock *FetchRepositoryMock) GetByIDCalls() []struct {
	C     context.Context
	Table string
	ID    string
	Item  interface{}
	Crit  *RepoCriterias
} {
	var calls []struct {
		C     context.Context
		Table string
		ID    string
		Item  interface{}
		Crit  *RepoCriterias
	}
	lockFetchRepositoryMockGetByID.RLock()
	calls = mock.calls.GetByID
	lockFetchRepositoryMockGetByID.RUnlock()
	return calls
}

// GetByRandom calls GetByRandomFunc.
func (mock *FetchRepositoryMock) GetByRandom(c context.Context, table string, item interface{}, crit *RepoCriterias) error {
	if mock.GetByRandomFunc == nil {
		panic("FetchRepositoryMock.GetByRandomFunc: method is nil but FetchRepository.GetByRandom was just called")
	}
	callInfo := struct {
		C     context.Context
		Table string
		Item  interface{}
		Crit  *RepoCriterias
	}{
		C:     c,
		Table: table,
		Item:  item,
		Crit:  crit,
	}
	lockFetchRepositoryMockGetByRandom.Lock()
	mock.calls.GetByRandom = append(mock.calls.GetByRandom, callInfo)
	lockFetchRepositoryMockGetByRandom.Unlock()
	return mock.GetByRandomFunc(c, table, item, crit)
}

// GetByRandomCalls gets all the calls that were made to GetByRandom.
// Check the length with:
//     len(mockedFetchRepository.GetByRandomCalls())
func (mock *FetchRepositoryMock) GetByRandomCalls() []struct {
	C     context.Context
	Table string
	Item  interface{}
	Crit  *RepoCriterias
} {
	var calls []struct {
		C     context.Context
		Table string
		Item  interface{}
		Crit  *RepoCriterias
	}
	lockFetchRepositoryMockGetByRandom.RLock()
	calls = mock.calls.GetByRandom
	lockFetchRepositoryMockGetByRandom.RUnlock()
	return calls
}