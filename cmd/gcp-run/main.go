package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog"

	"github.com/agilefoxHQ/go-function-template/config"
)

var (
	conf   config.Configuration
	logger zerolog.Logger
)

func init() {
	log.SetFlags(0)
	conf = config.Load()

	//ctx := context.Background()

	logger = zerolog.New(os.Stderr).With().Timestamp().Logger()

	if conf.Env != "development" {
		// Service level logging using cloud logging client. This client initializes with the right labels, resource
		// types and revision for cloud run deployed oldHandler
		//logger = zerolog.New(loggingWriter).Level(zerolog.InfoLevel)
	}
}

func serve(ctx context.Context) (err error) {
	// load the store and start services
	return
}

// main is maintained as a separate command package. It bootstraps the service.
// No other package should include the main package. Also, main should not use the external `services` directly.
func main() {
	// gracefully teardown the app on SIGINT / SIGTERM. For this we create a channel to listen for these signals.
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	// create a cancellable context - this context is what will control the server and store tear downs
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		// wait for the signal in a go routine
		oscall := <-sig
		logger.Info().Msgf("system call:%+v", oscall)
		// cancel the context when you receive this signal
		cancel()
	}()

	// Run serve with this cancellable context
	// all the application start logic can now be extracted away in serve, away from the main function
	// note to self:
	//	* the httpServer.ListenAndServe method unblocks immediately when httpServer.Shutdown is called
	//	* never return from the main function until you are actually ready to quit
	//	* an alternate approach would be to use synchronization primitives (wait groups, channels)
	if err := serve(ctx); err != nil {
		logger.Error().Err(err).Msg("failed to serve")
	}
}
