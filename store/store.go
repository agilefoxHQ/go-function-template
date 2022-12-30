package store

import (
	"context"
	"fmt"
	"os"

	"cloud.google.com/go/firestore"
	"github.com/rs/zerolog"

	"github.com/agilefoxHQ/go-function-template/config"
)

type Store struct {
	fs *firestore.Client
}

func LoadStore(ctx context.Context, c config.Configuration, ch <-chan os.Signal, logger *zerolog.Logger) (
	*Store, error,
) {
	fs, fsErr := NewFirestoreClient(ctx, c.ProjectID)
	if fsErr != nil {
		return nil, fmt.Errorf("firestore: %w", fsErr)
	}

	// keep a go routine running, listening to shut down events
	go func() {
		<-ch
		if err := fs.Close(); err != nil {
			logger.Error().Err(err).Msg("could not close firestore connection")
		}
		logger.Info().Msg("firestore gracefully shut down")
	}()
	logger.Info().Msg("Store loaded: client connections established")
	return &Store{fs: fs}, nil
}

// NewFirestoreClient creates a firestore client that could be used by the app
func NewFirestoreClient(ctx context.Context, projectID string) (
	*firestore.Client, error,
) {

	firestoreClient, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		return nil, fmt.Errorf("could not instantiate a Firestore Client: %w", err)
	}

	return firestoreClient, nil
}
