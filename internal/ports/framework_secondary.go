package ports

type DbPort interface {
	CloseDbConnection()
	AddToHistory(res int32, opearation string) error
}
