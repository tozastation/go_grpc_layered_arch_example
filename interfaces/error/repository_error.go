package error

import (
	"fmt"
	"github.com/tozastation/go-grpc-ddd-example/interfaces/error/irepoerror"
)

// RepositoryError  is .
type RepositoryError struct{}

// NewRepositoryError is .
func NewRepositoryError() irepoerror.IRepositoryError {
	return &RepositoryError{}
}

// SelectFailed is .
func (repoerr *RepositoryError) SelectFailed(err error) error {
	return fmt.Errorf("select query failed: %v", err)
}

// PrepareStatementFailed is .
func (repoerr *RepositoryError) PrepareStatementFailed(err error) error {
	return fmt.Errorf("prepare statement failed: %v", err)
}

// ExecuteQueryFailed is .
func (repoerr *RepositoryError) ExecuteQueryFailed(err error) error {
	return fmt.Errorf("execute query failed: %v", err)
}
