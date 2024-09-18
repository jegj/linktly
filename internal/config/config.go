package config

import (
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	LogLevel            string        `default:"WARN" envconfig:"LOG_LEVEL"`
	DBUser              string        `required:"true" envconfig:"DBUSER"`
	DBHost              string        `required:"true" envconfig:"DBHOST"`
	JwtSecret           string        `required:"true" envconfig:"JWT_SECRET"`
	DBName              string        `required:"true" envconfig:"DBNAME"`
	DBPasword           string        `required:"true" envconfig:"DBPASSWORD"`
	ServerAddress       string        `default:"8080" envconfig:"SERVER_ADDRESS"`
	ShutdownGracePeriod time.Duration `default:"30s" envconfig:"SHUTDOWN_GRACE_PERIOD"`
	ReadTimeout         time.Duration `default:"10s" envconfig:"READ_TIMEOUT"`
	IdleTimeout         time.Duration `default:"30s" envconfig:"IDLE_TIMEOUT"`
	WriteTimeout        time.Duration `default:"10s" envconfig:"WRITE_TIMEOUT"`
	DBPort              int           `required:"true" envconfig:"DBPORT"`
	LogHttpRequest      bool          `default:"false" envconfig:"LOG_HTTP_REQUEST"`
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

func NewConfig() Config {
	var env Config
	err := envconfig.Process("linktly", &env)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	return env
}
