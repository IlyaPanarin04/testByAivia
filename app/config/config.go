package config

type Datastore struct {
	Postgres string
}

type Config struct {
	Datastore Datastore
}

func NewConfig() Config {
	return Config{
		Datastore: Datastore{
			Postgres: "postgresql://postgres:test@localhost:5432/postgres",
		},
	}
}
