package services

import (
	"encoding/json"
	"os"

	consul "github.com/hashicorp/consul/api"
)

// AppConfig is for the inital config obtained from Consul
type AppConfig struct {
	Redis       RedisConfig
	Database    PostgresConfig
	Nats        NatsConfig
	Environment Environment
}

// Environment presents the env
type Environment struct {
	Type     string `json:"type"`
	LogLevel string `json:"log_level`
}

// RedisConfig is the configuration file for Redis
type RedisConfig struct {
	Host     string `json:"redis_host"`
	Username string `json:"user"`
	Password string `json:"password"`
}

// PostgresConfig is the configuration for DB
type PostgresConfig struct {
	Host     string `json:"db_host"`
	Username string `json:"db_user"`
	Password string `json:"db_password"`
	Database string `json:"db_name"`
	UseSSL   string `json:"db_usessl"`
	Port     int    `json:"db_port"`
}

// NatsConfig is the config for the Nats message bus
type NatsConfig struct {
	Host string `json:"nats_host"`
}

const (
	Production  string = "production"
	Development string = "development"
)

// ApplicationConfig is the app wide config
var ApplicationConfig *AppConfig

func configureDB(kv *consul.KV) PostgresConfig {
	// Lookup db config
	pair, _, err := kv.Get("DB_CONFIG", nil)
	if err != nil {
		panic(err)
	}
	var dbConfig PostgresConfig
	if err := json.Unmarshal(pair.Value, &dbConfig); err != nil {
		panic(err)
	}
	return dbConfig
}

func configureRedis(kv *consul.KV) RedisConfig {
	// Lookup redis config
	pair, _, err := kv.Get("REDIS_CONFIG", nil)
	if err != nil {
		panic(err)
	}
	var redisConfig RedisConfig
	if err := json.Unmarshal(pair.Value, &redisConfig); err != nil {
		panic(err)
	}
	return redisConfig
}

func configureNats(kv *consul.KV) NatsConfig {
	/// Lookup nats config
	pair, _, err := kv.Get("NATS_CONFIG", nil)
	if err != nil {
		panic(err)
	}
	var natsConfig NatsConfig
	if err := json.Unmarshal(pair.Value, &natsConfig); err != nil {
		panic(err)
	}
	return natsConfig
}

func getEnvironment(kv *consul.KV) Environment {
	/// Finally look up environment
	pair, _, err := kv.Get("ENVIRONMENT", nil)
	if err != nil {
		panic(err)
	}
	var environment Environment
	if err := json.Unmarshal(pair.Value, &environment); err != nil {
		panic(err)
	}
	return environment
}

func init() {
	// Get a new client
	Logger.Info("Starting config bootstrap")
	consulHost := os.Getenv("consulhost")
	config := consul.Config{
		Address: consulHost,
	}
	client, err := consul.NewClient(&config)
	if err != nil {
		panic(err)
	}

	// Get a handle to the KV API
	kv := client.KV()
	env := getEnvironment(kv)

	ApplicationConfig = &AppConfig{
		Redis:       configureRedis(kv),
		Database:    configureDB(kv),
		Nats:        configureNats(kv),
		Environment: env,
	}
}
