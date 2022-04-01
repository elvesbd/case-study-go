package ports

type DbPort interface {
	CloseDbConnection()
	addToHistory(answer int32, operation string) error
}
