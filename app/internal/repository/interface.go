package repository

type DataStorePostgres interface {
	TestConnection()
}
