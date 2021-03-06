// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package contract

import (
	"context"
	"sync"
)

var (
	lockAggregationRepositoryMockCountByCriteria sync.RWMutex
	lockAggregationRepositoryMockCountUnique     sync.RWMutex
	lockAggregationRepositoryMockSumByCriteria   sync.RWMutex
)

// Ensure, that AggregationRepositoryMock does implement AggregationRepository.
// If this is not the case, regenerate this file with moq.
var _ AggregationRepository = &AggregationRepositoryMock{}

// AggregationRepositoryMock is a mock implementation of AggregationRepository.
//
//     func TestSomethingThatUsesAggregationRepository(t *testing.T) {
//
//         // make and configure a mocked AggregationRepository
//         mockedAggregationRepository := &AggregationRepositoryMock{
//             CountByCriteriaFunc: func(c context.Context, table string, crit *RepoCriterias) (uint64, error) {
// 	               panic("mock out the CountByCriteria method")
//             },
//             CountUniqueFunc: func(c context.Context, table string, field string, name string) (uint64, error) {
// 	               panic("mock out the CountUnique method")
//             },
//             SumByCriteriaFunc: func(c context.Context, table string, field string, crit *RepoCriterias) (uint64, error) {
// 	               panic("mock out the SumByCriteria method")
//             },
//         }
//
//         // use mockedAggregationRepository in code that requires AggregationRepository
//         // and then make assertions.
//
//     }
type AggregationRepositoryMock struct {
	// CountByCriteriaFunc mocks the CountByCriteria method.
	CountByCriteriaFunc func(c context.Context, table string, crit *RepoCriterias) (uint64, error)

	// CountUniqueFunc mocks the CountUnique method.
	CountUniqueFunc func(c context.Context, table string, field string, name string) (uint64, error)

	// SumByCriteriaFunc mocks the SumByCriteria method.
	SumByCriteriaFunc func(c context.Context, table string, field string, crit *RepoCriterias) (uint64, error)

	// calls tracks calls to the methods.
	calls struct {
		// CountByCriteria holds details about calls to the CountByCriteria method.
		CountByCriteria []struct {
			// C is the c argument value.
			C context.Context
			// Table is the table argument value.
			Table string
			// Crit is the crit argument value.
			Crit *RepoCriterias
		}
		// CountUnique holds details about calls to the CountUnique method.
		CountUnique []struct {
			// C is the c argument value.
			C context.Context
			// Table is the table argument value.
			Table string
			// Field is the field argument value.
			Field string
			// Name is the name argument value.
			Name string
		}
		// SumByCriteria holds details about calls to the SumByCriteria method.
		SumByCriteria []struct {
			// C is the c argument value.
			C context.Context
			// Table is the table argument value.
			Table string
			// Field is the field argument value.
			Field string
			// Crit is the crit argument value.
			Crit *RepoCriterias
		}
	}
}

// CountByCriteria calls CountByCriteriaFunc.
func (mock *AggregationRepositoryMock) CountByCriteria(c context.Context, table string, crit *RepoCriterias) (uint64, error) {
	if mock.CountByCriteriaFunc == nil {
		panic("AggregationRepositoryMock.CountByCriteriaFunc: method is nil but AggregationRepository.CountByCriteria was just called")
	}
	callInfo := struct {
		C     context.Context
		Table string
		Crit  *RepoCriterias
	}{
		C:     c,
		Table: table,
		Crit:  crit,
	}
	lockAggregationRepositoryMockCountByCriteria.Lock()
	mock.calls.CountByCriteria = append(mock.calls.CountByCriteria, callInfo)
	lockAggregationRepositoryMockCountByCriteria.Unlock()
	return mock.CountByCriteriaFunc(c, table, crit)
}

// CountByCriteriaCalls gets all the calls that were made to CountByCriteria.
// Check the length with:
//     len(mockedAggregationRepository.CountByCriteriaCalls())
func (mock *AggregationRepositoryMock) CountByCriteriaCalls() []struct {
	C     context.Context
	Table string
	Crit  *RepoCriterias
} {
	var calls []struct {
		C     context.Context
		Table string
		Crit  *RepoCriterias
	}
	lockAggregationRepositoryMockCountByCriteria.RLock()
	calls = mock.calls.CountByCriteria
	lockAggregationRepositoryMockCountByCriteria.RUnlock()
	return calls
}

// CountUnique calls CountUniqueFunc.
func (mock *AggregationRepositoryMock) CountUnique(c context.Context, table string, field string, name string) (uint64, error) {
	if mock.CountUniqueFunc == nil {
		panic("AggregationRepositoryMock.CountUniqueFunc: method is nil but AggregationRepository.CountUnique was just called")
	}
	callInfo := struct {
		C     context.Context
		Table string
		Field string
		Name  string
	}{
		C:     c,
		Table: table,
		Field: field,
		Name:  name,
	}
	lockAggregationRepositoryMockCountUnique.Lock()
	mock.calls.CountUnique = append(mock.calls.CountUnique, callInfo)
	lockAggregationRepositoryMockCountUnique.Unlock()
	return mock.CountUniqueFunc(c, table, field, name)
}

// CountUniqueCalls gets all the calls that were made to CountUnique.
// Check the length with:
//     len(mockedAggregationRepository.CountUniqueCalls())
func (mock *AggregationRepositoryMock) CountUniqueCalls() []struct {
	C     context.Context
	Table string
	Field string
	Name  string
} {
	var calls []struct {
		C     context.Context
		Table string
		Field string
		Name  string
	}
	lockAggregationRepositoryMockCountUnique.RLock()
	calls = mock.calls.CountUnique
	lockAggregationRepositoryMockCountUnique.RUnlock()
	return calls
}

// SumByCriteria calls SumByCriteriaFunc.
func (mock *AggregationRepositoryMock) SumByCriteria(c context.Context, table string, field string, crit *RepoCriterias) (uint64, error) {
	if mock.SumByCriteriaFunc == nil {
		panic("AggregationRepositoryMock.SumByCriteriaFunc: method is nil but AggregationRepository.SumByCriteria was just called")
	}
	callInfo := struct {
		C     context.Context
		Table string
		Field string
		Crit  *RepoCriterias
	}{
		C:     c,
		Table: table,
		Field: field,
		Crit:  crit,
	}
	lockAggregationRepositoryMockSumByCriteria.Lock()
	mock.calls.SumByCriteria = append(mock.calls.SumByCriteria, callInfo)
	lockAggregationRepositoryMockSumByCriteria.Unlock()
	return mock.SumByCriteriaFunc(c, table, field, crit)
}

// SumByCriteriaCalls gets all the calls that were made to SumByCriteria.
// Check the length with:
//     len(mockedAggregationRepository.SumByCriteriaCalls())
func (mock *AggregationRepositoryMock) SumByCriteriaCalls() []struct {
	C     context.Context
	Table string
	Field string
	Crit  *RepoCriterias
} {
	var calls []struct {
		C     context.Context
		Table string
		Field string
		Crit  *RepoCriterias
	}
	lockAggregationRepositoryMockSumByCriteria.RLock()
	calls = mock.calls.SumByCriteria
	lockAggregationRepositoryMockSumByCriteria.RUnlock()
	return calls
}
