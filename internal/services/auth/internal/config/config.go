package config

type Config struct {
	Transport struct {
		Rpc struct {
			Addres string
		}
	}
	Databases struct {
		AuthDb PostgresDB
	}
}

type PostgresDB struct {
	DSN string
}
