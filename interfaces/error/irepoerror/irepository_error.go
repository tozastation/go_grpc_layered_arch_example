package irepoerror

// IRepositoryError is error interface for repository
type IRepositoryError interface {
	SelectFailed(err error) error
	PrepareStatementFailed(err error) error
	ExecuteQueryFailed(err error) error
}
