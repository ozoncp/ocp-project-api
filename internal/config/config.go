package config

import (
	"fmt"
	"os"

	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Service  ServiceConfig  `yaml:"service"`
	DB       PostgresConfig `yaml:"storage"`
	Producer KafkaConfig    `yaml:"kafka"`
	Metrics  MetricsConfig  `yaml:"metrics"`
}

type ServiceConfig struct {
	Host     string `yaml:"host"`
	GrpcPort string `yaml:"grpc_port"`
	HttpPort string `yaml:"http_port"`
}

type KafkaConfig struct {
	Brokers     []string `yaml:"brokers"`
	EventsTopic string   `yaml:"events_topic"`
	PingTopic   string   `yaml:"ping_topic"`
	Capacity    uint64   `yaml:"capacity"`
}

type PostgresConfig struct {
	Driver    string `yaml:"driver"`
	User      string `yaml:"user"`
	Name      string `yaml:"db_name"`
	Sslmode   string `yaml:"ssl_mode"`
	ChunkSize int    `yaml:"chunk_size"`
}

type MetricsConfig struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

var Global = Config{
	Service: ServiceConfig{
		Host:     "0.0.0.0",
		GrpcPort: "8082",
		HttpPort: "8080",
	},
	Producer: KafkaConfig{
		Brokers:     []string{"127.0.0.1:9094"},
		EventsTopic: "events",
		PingTopic:   "ping",
		Capacity:    256,
	},
	DB: PostgresConfig{
		Driver:    "postgres",
		User:      "lobanov",
		Name:      "ocp",
		Sslmode:   "disable",
		ChunkSize: 1,
	},
	Metrics: MetricsConfig{
		Host: "0.0.0.0",
		Port: "9010",
	},
}

func (c ServiceConfig) GrpcEndpoint() string {
	return getEndpoint(c.Host, c.GrpcPort)
}

func (c ServiceConfig) HttpEndpoint() string {
	return getEndpoint(c.Host, c.HttpPort)
}

func (c MetricsConfig) Endpoint() string {
	return getEndpoint(c.Host, c.Port)
}

func getEndpoint(host string, port string) string {
	return fmt.Sprintf("%s:%s", host, port)
}

func LoadGlobal(path string) {
	if err := validateConfigPath(path); err != nil {
		log.Error().Msgf("Validation config path error: %v", err)
		return
	}
	// Create config structure
	config := Global

	// Open config file
	file, err := os.Open(path)
	if err != nil {
		log.Error().Msgf("Opening config path error: %v", err)
		return
	}
	defer file.Close()

	// Init new YAML decode
	d := yaml.NewDecoder(file)
	d.SetStrict(true)

	// Start YAML decoding from file
	if err := d.Decode(&config); err != nil {
		log.Error().Msgf("Decoding config error: %v", err)
		return
	}

	Global = config
}

// ValidateConfigPath just makes sure, that the path provided is a file,
// that can be read
func validateConfigPath(path string) error {
	s, err := os.Stat(path)
	if err != nil {
		return err
	}
	if s.IsDir() {
		return fmt.Errorf("'%s' is a directory, not a normal file", path)
	}
	return nil
}
