package config

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"os"
	"time"
)

const (
	defaultConfigPath = "./configs/config.yaml"
	defaultEnvPath    = ".env"
)

type (
	Config struct {
		App  `yaml:"app"`
		GRPC `yaml:"http"`
		PG   `yaml:"pg"`
		Log  `yaml:"logger"`
	}

	App struct {
		Name    string `env:"APP_NAME"            env-default:"SneakerStore-UserService" yaml:"name"`
		Version string `env:"APP_VERSION"         env-default:"v1.0.0"         yaml:"version"`
	}

	GRPC struct {
		Port    string        `env:"GRPC_PORT"    env-default:":8080" yaml:"port"`
		Timeout time.Duration `env:"GRPC_TIMEOUT" env-default:"5s"    yaml:"timeout"`
	}

	PG struct {
		PoolMax int    `env:"PG_POOL_MAX" env-default:"2"     yaml:"poolMax"`
		URL     string `env:"PG_URL"      env-required:"true" yaml:"url"`
	}

	Log struct {
		Level string `env:"LOG_LEVEL" env-default:"debug" yaml:"logLevel"`
	}
)

func MustLoad() *Config {
	configPath := fetchConfigPath()
	if configPath == "" {
		panic("config path is empty")
	}

	return MustLoadPath(configPath, defaultEnvPath)
}

func MustLoadPath(configPath, envPath string) *Config {
	// check if file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("config file does not exist: " + configPath)
	}

	var cfg Config

	_ = godotenv.Load(envPath)

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic("cannot read config: " + err.Error())
	}

	return &cfg
}

func fetchConfigPath() string {
	var res string

	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	if res == "" {
		res = defaultConfigPath
	}

	return res
}
