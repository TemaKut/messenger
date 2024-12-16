package config

func DefaultConfig() *Config {
	var c Config

	c.State.Environment = EnvironmentLocal

	c.State.Transport.Rpc.Addres = ":8001"

	c.State.Databases.AuthDb.DSN = "postgres://root:root@localhost:5432/auth?sslmode=disable"

	return &c
}
