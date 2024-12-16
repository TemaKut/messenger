package config

type Environment int

const (
	EnvironmentLocal = iota
	EnvironmentStage
	EnvironmentProd
)

type Config struct {
	State ConfigState
}

func (c *Config) GetState() ConfigState {
	// TODO тут динамический контекст из consul
	return c.State
}

type ConfigState struct {
	Debug       bool
	Environment Environment
	Transport   struct {
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
