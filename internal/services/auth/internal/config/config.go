package config

type Environment int

const (
	EnvironmentLocal = iota
	EnvironmentStage
	EnvironmentProd
)

type Config struct {
	state ConfigState
}

func (c *Config) GetState() ConfigState {
	// TODO тут динамический контекст из consul
	return c.state
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

	Broker struct {
		Addrs []string
	}
}

type PostgresDB struct {
	DSN string
}
