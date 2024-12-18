package kafka

type Environment string

var (
	EnvironmentLocal Environment = "local"
	EnvironmentStage             = "stage"
	EnvironmentProd              = "prod"
)

type TopicBuilder struct {
	env Environment
}

func NewTopicBuilder(env Environment) *TopicBuilder {
	return &TopicBuilder{env: env}
}

func (t *TopicBuilder) UserCreatedV1Topic() string {
	return string(t.env) + ".user_created.v1"
}
