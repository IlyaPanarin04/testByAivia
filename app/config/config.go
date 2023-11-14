package config

type Datastore struct {
	Postgres string
}

type Config struct {
	Port      string
	Datastore Datastore
}

func NewConfig() Config {
	return Config{
		Port: ":4000",
		Datastore: Datastore{
			Postgres: "postgresql://postgres:test@localhost:5432/postgres",
		},
	}
}
