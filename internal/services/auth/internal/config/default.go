package config

func DefaultConfig() *Config {
	var c Config

	c.state.Environment = EnvironmentLocal

	c.state.Transport.Rpc.Addres = ":8001"

	c.state.Databases.AuthDb.DSN = "postgres://root:root@localhost:5432/auth?sslmode=disable"

	c.state.Broker.Addrs = []string{"localhost:9092"}

	return &c
}
