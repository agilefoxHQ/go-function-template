package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/rs/zerolog"

	"github.com/agilefoxHQ/go-function-template/config"
	"github.com/agilefoxHQ/go-function-template/handlers"
	"github.com/agilefoxHQ/go-function-template/store"
)

var (
	c       config.Configuration
	logger  zerolog.Logger
	st      *store.Store
	handler *http.ServeMux
)

// init is used to initialise the package to optimise for hot and cold restarts.
// With all the hate around init in go, I think this use case is the one where init really shines.
func init() {
	c = config.Load()
	var err error
	logger = zerolog.New(os.Stderr).With().Timestamp().Logger()
	ctx, cancel := createContext(c.Timeout)
	defer cancel()
	// gracefully close the store on os.Interrupt
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	st, err = store.LoadStore(ctx, c, sig, &logger)
	if err != nil {
		logger.Fatal().Msgf("could not instantiate store: %v", err)
	}
	handler = handlers.NewHandler(st, &logger)
}

func main() {
	logger.Info().Msgf("\n%s listening on port :%s", c.FunctionName, c.Port)
	logger.Fatal().Err(http.ListenAndServe(fmt.Sprintf(":%s", c.Port), handler))
}

// Use context.Background() because the underlying container will persist across invocations on the GCP serverless infra
// The function runtime comes with its own timeout, but I would like to gracefully handle store shutdowns,
// if any by using my own timeout and listening to context cancellations and SIG_INT signals
func createContext(svcTimeout time.Duration) (context.Context, context.CancelFunc) {
	d := time.Now().Add(svcTimeout)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	return ctx, cancel
}
