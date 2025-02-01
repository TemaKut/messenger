package config

var DefaultConfig = Config{}

func init() {
	DefaultConfig.Debug = true
	DefaultConfig.Environment = EnvironmentLocal

	DefaultConfig.Server.Websocket.Addr = ":8001"

	DefaultConfig.Clients.AuthService.Addr = ":8002"
}
