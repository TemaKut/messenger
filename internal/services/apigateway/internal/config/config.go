package config

type Environment int

const (
	EnvironmentLocal Environment = iota
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
		Broker struct {
			Addrs []string
		}
		Websocket struct {
			Addr string
		}
	}
}

type PostgresDB struct {
	DSN string
}
