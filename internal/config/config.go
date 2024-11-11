package config

import (
	"fmt"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	App     AppConfig     `yaml:"app"`
	HTTPSrv HTTPSrvConfig `yaml:"http"`
	JWT     JWTConfig     `yaml:"jwt"`

	Database DBConfig    `yaml:"postgres"`
	S3       S3Config    `yaml:"minio"`
	Redis    RedisConfig `yaml:"redis"`
}

type AppConfig struct {
	Name    string `yaml:"name" env-required:"true"`
	Version int    `yaml:"version" env-required:"true"`
	Env     string `yaml:"env" env-required:"true"`
}

type HTTPSrvConfig struct {
	Port        int           `yaml:"port" env-required:"true"`
	Timeout     time.Duration `yaml:"timeout" env-required:"true"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-required:"true"`
}

type JWTConfig struct {
	TokenTTL time.Duration `yaml:"token_ttl" env-required:"true"`
	SignKey  string        `yaml:"sign_key" env-required:"true"`
}

type DBConfig struct {
	Host     string `yaml:"host" env-required:"true"`
	Port     int    `yaml:"port" env-required:"true"`
	Name     string `yaml:"name" env-required:"true"`
	User     string `yaml:"user" env-required:"true"`
	Password string `yaml:"password" env-required:"true"`
}

type S3Config struct {
	Bucket    string `yaml:"bucket" env-required:"true"`
	Host      string `yaml:"host" env-required:"true"`
	PublicURL string `yaml:"public_url" env-required:"true"`
	AcessKey  string `yaml:"access_key" env-required:"true"`
	SecretKey string `yaml:"secret_key" env-required:"true"`
	Region    string `yaml:"region" env-required:"true"`
	UseSSL    *bool  `yaml:"ssl" env-required:"true"`
}

type RedisConfig struct {
	Host string `yaml:"host" env-required:"true"`
	Port int    `yaml:"port" env-required:"true"`
	DB   *int   `yaml:"db" env-required:"true"`

	ApartmentTTL time.Duration `yaml:"apartment_ttl" env-required:"true"`
}

func Load(path string) (*Config, error) {
	var cfg Config
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, fmt.Errorf("config file by path(%s) not found", path)
	}

	err := cleanenv.ReadConfig(path, &cfg)
	if err != nil {
		return nil, fmt.Errorf("read config: %w", err)
	}
	return &cfg, nil
}
