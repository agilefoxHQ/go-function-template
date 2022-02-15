package go_function_template

import (
	"log"
	"os"
)

// GOOGLE_CLOUD_PROJECT is automatically set by the Cloud Functions runtime.
var projectID = os.Getenv("GOOGLE_CLOUD_PROJECT")

// init is used to initialise the package to optimise for hot and cold restarts. With all
// the hate around init in go, I think this use case is the one where init really shines.
func init() {
	// Use context.Background() because the app/client should persist across
	// invocations.
	// ctx := context.Background()

	// You can have common logic here like database connections with pooling
	// enabled or logging clients. Refer to `cloud.google.com/go/logging`
	// for a function level logging using cloud logging client.
	// This client initializes with the right labels,
	// resource types and revisions for cloud functions saving a lot of bootstrapping code
	log.SetFlags(0)

}
