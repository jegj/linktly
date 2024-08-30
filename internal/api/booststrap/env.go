package bootstrap

import (
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type EnvVar struct {
	ServerAddress       string        `default:"8080" envconfig:"SERVER_ADDRESS"`
	DBUser              string        `required:"true" envconfig:"DBUSER"`
	DBHost              string        `required:"true" envconfig:"DBHOST"`
	DBPort              int           `required:"true" envconfig:"DBPORT"`
	DBName              string        `required:"true" envconfig:"DBNAME"`
	DBPasword           string        `required:"true" envconfig:"DBPASSWORD"`
	ShutdownGracePeriod time.Duration `default:"30s" envconfig:"SHUTDOWN_GRACE_PERIOD"`
	WriteTimeout        time.Duration `default:"10s" envconfig:"WRITE_TIMEOUT"`
	ReadTimeout         time.Duration `default:"10s" envconfig:"READ_TIMEOUT"`
	IdleTimeout         time.Duration `default:"30s" envconfig:"IDLE_TIMEOUT"`
	LogHttpRequest      bool          `default:"false" envconfig:"LOG_HTTP_REQUEST"`
	LogLevel            string        `default:"WARN" envconfig:"LOG_LEVEL"`
}

func (envVar EnvVar) GetDBConnectionString() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%v/%s",
		envVar.DBUser,
		envVar.DBPasword,
		envVar.DBHost,
		envVar.DBPort,
		envVar.DBName,
	)
}

func NewEnvVar() *EnvVar {
	var env EnvVar
	err := envconfig.Process("linktly", &env)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	return &env
}
