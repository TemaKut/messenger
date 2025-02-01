package config

var DefaultConfig = Config{}

func init() {
	DefaultConfig.Debug = true
	DefaultConfig.Environment = EnvironmentLocal

	DefaultConfig.Server.GRPC.Addr = ":8002"
}
