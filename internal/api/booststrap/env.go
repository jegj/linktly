package bootstrap

import (
	"log"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type EnvVar struct {
	ServerAddress       string        `default:"8080" envconfig:"SERVER_ADDRESS"`
	PostgresUser        string        `required:"true" envconfig:"POSTGRES_USER"`
	PostgresDatabase    string        `required:"true" envconfig:"POSTGRES_DB"`
	PostgresPassword    string        `required:"true" envconfig:"POSTGRES_PASSWORD"`
	ShutdownGracePeriod time.Duration `default:"30s" envconfig:"SHUTDOWN_GRACE_PERIOD"`
	WriteTimeout        time.Duration `default:"10s" envconfig:"WRITE_TIMEOUT"`
	ReadTimeout         time.Duration `default:"10s" envconfig:"READ_TIMEOUT"`
	IdleTimeout         time.Duration `default:"30s" envconfig:"IDLE_TIMEOUT"`
}

func NewEnvVar() *EnvVar {
	var env EnvVar
	err := envconfig.Process("linktly", &env)
	if err != nil {
		log.Fatal(err.Error())
	}
	return &env
}
