package config

// TODO struct теги с .env или consul
type Config struct {
	Debug       bool
	Environment Environment

	Server struct {
		Websocket struct {
			Addr string
		}
	}

	Clients struct {
		AuthService struct {
			Addr string
		}
	}
}

type Environment string

const (
	EnvironmentLocal = "LOCAL"
	EnvironmentStage = "STAGE"
)

func NewConfig() *Config {
	return &DefaultConfig // TODO В дефолтный конфиг парсим .env
}
