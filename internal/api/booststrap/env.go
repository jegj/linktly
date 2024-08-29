package bootstrap

import (
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/kelseyhightower/envconfig"
)

// FIXME: ENV VAR  FOR POSTGRES
type EnvVar struct {
	ServerAddress       string        `default:"8080" envconfig:"SERVER_ADDRESS"`
	PostgresUser        string        `required:"true" envconfig:"POSTGRES_USER"`
	PostgresHost        string        `required:"true" envconfig:"PGHOST"`
	PostgresPort        int           `required:"true" envconfig:"PGPORT"`
	PostgresDatabase    string        `required:"true" envconfig:"POSTGRES_DB"`
	PostgresPassword    string        `required:"true" envconfig:"POSTGRES_PASSWORD"`
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
		envVar.PostgresUser,
		envVar.PostgresPassword,
		envVar.PostgresHost,
		envVar.PostgresPort,
		envVar.PostgresDatabase,
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
