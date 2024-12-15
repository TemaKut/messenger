package config

func DefaultConfig() *Config {
	var c Config

	c.Transport.Rpc.Addres = ":8001"

	c.Databases.AuthDb.DSN = "postgres://root:root@localhost:5432/auth"

	return &c
}
