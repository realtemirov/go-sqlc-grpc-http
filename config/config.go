package config

import (
	"embed"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/caarlos0/env/v9"
	"github.com/joho/godotenv"
	"github.com/realtemirov/go-sqlc-grpc-http/pkg/logger"
	"gopkg.in/yaml.v3"
)

type AppMode string

const (
	DEVELOPMENT AppMode = "DEVELOPMENT"
	PRODUCTION  AppMode = "PRODUCTION"
)

//go:embed configs
var configs embed.FS

type Config struct {
	Logging logger.LoggingConfig `yaml:"logging"`
	Mode    string               `env:"APPLICATION_MODE" envDefault:"development" yaml:"mode"`

	Project struct {
		Name           string        `env:"PROJECT_NAME" yaml:"name"`
		Version        string        `env:"PROJECT_VERSION" yaml:"version"`
		Timeout        time.Duration `env:"PROJECT_TIMEOUT" yaml:"timeout"`
		SwaggerEnabled bool          `env:"PROJECT_SWAGGER_ENABLED" yaml:"swagger_enabled"`
		CacheTimeout   time.Duration `env:"PROJECT_CACHE_TIMEOUT" yaml:"cache_timeout"`
	} `yaml:"project"`

	HTTP struct {
		Host string `env:"HTTP_HOST" yaml:"host"`
		Port string `env:"PORT" yaml:"port"`

		URL string `env:"HTTP_URL" yaml:"url"`
	} `yaml:"http"`

	Grpc struct {
		Host string `env:"GRPC_HOST" yaml:"host"`
		Port string `env:"GRPC_PORT" yaml:"port"`

		URL string `env:"GRPC_URL" yaml:"url"`
	} `yaml:"grpc"`

	PSQL struct {
		URI string `env:"PSQL_URI" yaml:"uri"`
	} `yaml:"psql"`

	StorageConfig struct {
		URI                  string        `env:"STORAGE_URI" yaml:"uri"`
		User                 string        `env:"STORAGE_MINIO_USER" yaml:"user"`
		Password             string        `env:"STORAGE_MINIO_PASSWORD" yaml:"password"`
		BucketName           string        `env:"STORAGE_BUCKET_NAME" yaml:"bucket_name"`
		Secure               bool          `env:"STORAGE_SECURE" yaml:"secure"`
		AccessTokenExpiresAt time.Duration `env:"STORAGE_ACCESS_TOKEN_EXPIRES_AT" envDefault:"2h" yaml:"token_expires_at"`
	} `yaml:"storage"`
}

func Load() *Config {
	var cfg Config
	err := godotenv.Load(".env")
	if err != nil && !os.IsNotExist(err) {
		log.Fatal(err)
	}

	configPath := getConfigPath(getAppMode())

	file, err := configs.ReadFile(configPath)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(file, &cfg)
	if err != nil {
		panic(err)
	}

	if errParse := env.Parse(&cfg); errParse != nil {
		log.Println(errParse.Error())
		panic("unmarshal from environment error")
	}

	cfg.MakeGrpcURL()
	cfg.MakeHTTPURL()

	return &cfg
}

func getAppMode() AppMode {
	mode := AppMode(os.Getenv("APPLICATION_MODE"))

	if mode != DEVELOPMENT {
		mode = PRODUCTION
	}

	return mode
}

func getConfigPath(appMode AppMode) string {
	suffix := "dev"
	if appMode == PRODUCTION {
		suffix = "prod"
	}

	return fmt.Sprintf("configs/%s.yaml", suffix)
}

func (c *Config) MakeGrpcURL() {
	c.Grpc.URL = fmt.Sprintf("%s:%s", c.Grpc.Host, c.Grpc.Port)
}

func (c *Config) MakeHTTPURL() {
	c.HTTP.URL = fmt.Sprintf("%s:%s", c.HTTP.Host, c.HTTP.Port)
}
