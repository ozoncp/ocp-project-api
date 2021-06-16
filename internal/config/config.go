package config

type Config struct {
	Producer KafkaConfig
	DB       PostgresConfig
}

type KafkaConfig struct {
	Brokers     []string
	EventsTopic string
	PingTopic   string
	Capacity    uint64
}

type PostgresConfig struct {
	Driver  string
	User    string
	Name    string
	Sslmode string
}

var Global = Config{
	Producer: KafkaConfig{
		Brokers:     []string{"127.0.0.1:9094"},
		EventsTopic: "events",
		PingTopic:   "ping",
		Capacity:    256,
	},
	DB: PostgresConfig{
		Driver:  "postgres",
		User:    "lobanov",
		Name:    "ocp",
		Sslmode: "disable",
	},
}
