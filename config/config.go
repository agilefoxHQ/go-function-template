package config

import (
	"log"
	"os"
	"time"

	"github.com/joeshaw/envdecode"
)

// Version is set during the compilation time to the git hash of the code version that the binary was build from.
// It must not be modified during the runtime. Check Makefile `go build` command to see this assignment.
var (
	Version string
	Tag     string
)

// Configuration is exposed for the application to be able to see what the application config looks like
type Configuration struct {
	// service config
	Env          string        `env:"ENVIRONMENT,default=development"`
	ServiceName  string        `env:"SERVICE_NAME,required"`
	ProjectID    string        `env:"GOOGLE_CLOUD_PROJECT,required"` // GCP function runtime
	Timeout      time.Duration `env:"TIMEOUT,default=10s,strict"`
	FunctionName string        `env:"FUNCTION_NAME,default=do_that_thing"`
	Port         string        `env:"PORT,default=8080"`
}

// Load loads the configuration using the envdecode package
func Load() Configuration {
	var c Configuration
	if err := envdecode.Decode(&c); err != nil {
		log.Printf("error reading env vars: %v\n", err)
		os.Exit(2)
	}
	return c
}
