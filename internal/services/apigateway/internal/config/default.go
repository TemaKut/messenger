package config

func DefaultConfig() *Config {
	var c Config

	c.state.Environment = EnvironmentLocal

	c.state.Transport.Broker.Addrs = []string{"localhost:9092"}

	c.state.Transport.Websocket.Addr = ":12001"

	return &c
}
