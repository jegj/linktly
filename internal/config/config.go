package config

import (
	"crypto/rsa"
	"encoding/base64"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	LinktlyPrivateKey       string        `required:"true" envconfig:"LINKTLY_PRIVATE_KEY"`
	DBUser                  string        `required:"true" envconfig:"DBUSER"`
	DBHost                  string        `required:"true" envconfig:"DBHOST"`
	DBName                  string        `required:"true" envconfig:"DBNAME"`
	DBPasword               string        `required:"true" envconfig:"DBPASSWORD"`
	ServerAddress           string        `default:"3000" envconfig:"SERVER_ADDRESS"`
	LogLevel                string        `default:"WARN" envconfig:"LOG_LEVEL"`
	LinktlyPublicKey        string        `required:"true" envconfig:"LINKTLY_PUBLIC_KEY"`
	ShutdownGracePeriod     time.Duration `default:"30s" envconfig:"SHUTDOWN_GRACE_PERIOD"`
	WriteTimeout            time.Duration `default:"10s" envconfig:"WRITE_TIMEOUT"`
	DBPort                  int           `required:"true" envconfig:"DBPORT"`
	IdleTimeout             time.Duration `default:"30s" envconfig:"IDLE_TIMEOUT"`
	ReadTimeout             time.Duration `default:"10s" envconfig:"READ_TIMEOUT"`
	AccessTokenExpTime      time.Duration `default:"1h" required:"true" envconfig:"ACCESS_TOKEN_EXP_TIME"`
	RefreshTokenExpTime     time.Duration `default:"120h" required:"true" envconfig:"REFRESH_TOKEN_EXP_TIME"`
	LogHttpRequest          bool          `default:"false" envconfig:"LOG_HTTP_REQUEST"`
	HTTPCookieSecure        bool          `default:"true" envconfig:"HTTP_COOKIE_SECURE"`
	PgPoolMaxConn           int32         `default:"10" envconfig:"PGPOOL_MAX_CONN"`
	PgPoolMinConn           int32         `default:"1" envconfig:"PGPOOL_MIN_CONN"`
	PgPoolConnLifeTime      time.Duration `default:"30m" envconfig:"PGPOOL_CONN_LIFETIME"`
	PgPoolMaxConnIdleTime   time.Duration `default:"10m" envconfig:"PGPOOL_MAX_CONN_IDLE_TIME"`
	PgPoolHealthCheckPeriod time.Duration `default:"1m" envconfig:"PGPOOL_HEALTHCHECK_PERIOD"`
}

func (envVar Config) GetDBConnectionString() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%v/%s",
		envVar.DBUser,
		envVar.DBPasword,
		envVar.DBHost,
		envVar.DBPort,
		envVar.DBName,
	)
}

func (envVar Config) GetPrivateKey() (*rsa.PrivateKey, error) {
	privateKeyBase64 := envVar.LinktlyPrivateKey
	privateKeyBytes, err := base64.StdEncoding.DecodeString(privateKeyBase64)
	if err != nil {
		return nil, err
	}
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyBytes)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}

func (envVar Config) GetPublicKey() (*rsa.PublicKey, error) {
	publicKeyBase64 := envVar.LinktlyPublicKey
	publicKeyBytes, err := base64.StdEncoding.DecodeString(publicKeyBase64)
	if err != nil {
		return nil, err
	}
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicKeyBytes)
	if err != nil {
		return nil, err
	}
	return publicKey, nil
}

func NewConfig() Config {
	var env Config
	err := envconfig.Process("linktly", &env)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	return env
}
