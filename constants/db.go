package constants

import "errors"

type dbError struct {
	ERROR_WRONG_TYPE       error
	ERROR_RECORD_NOT_FOUND error
}

var (
	ContextKeyTransaction = "Tx"
	DBError               = dbError{
		ERROR_WRONG_TYPE:       errors.New("unsupport_type"),
		ERROR_RECORD_NOT_FOUND: errors.New("record_not_found"),
	}
)
