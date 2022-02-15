package go_function_template

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/logging"
	myLogger "github.com/agilefoxHQ/pkg/gcp-zerolog"
	"github.com/rs/zerolog"

	"cloud.google.com/go/firestore"
)

// GOOGLE_CLOUD_PROJECT is automatically set by the Cloud Functions runtime.
var projectID = os.Getenv("GOOGLE_CLOUD_PROJECT")

// client is a Firestore client, reused between function invocations.
var client *firestore.Client

var logger zerolog.Logger

// init is used to initialise the package to optimise for hot and cold restarts. With all
// the hate around init in go, I think this use case is the one where init really shines.
func init() {

	log.SetFlags(0)

	// Use context.Background() because the app/client should persist across
	// invocations.
	ctx := context.Background()

	// Function level logging using cloud logging client. This client initializes with the
	// right labels, resource types and revisions for cloud functions
	lc, err := logging.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// sets the name of the log to write to.
	lcLogger := lc.Logger("create-user-log")
	loggingWriter, lcErr := myLogger.NewCloudLoggingWriter(
		ctx, myLogger.CloudLoggingOptions{
			SeverityMap: myLogger.DefaultSeverityMap,
			Logger:      lcLogger,
		},
	)
	if lcErr != nil {
		log.Panicf("could not create a CloudLoggingWriter: %v", lcErr)
	}

	logger = zerolog.New(loggingWriter).Level(zerolog.InfoLevel)

}
