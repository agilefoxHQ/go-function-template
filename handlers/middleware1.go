package handlers

import (
	"net/http"
)

func (h *handler) middleware1(handler http.Handler) http.Handler {
	// Call the underlying handler
	return http.HandlerFunc(
		func(writer http.ResponseWriter, request *http.Request) {
			// Find real closure to your problems, do something worthwhile in life
			handler.ServeHTTP(writer, request)
		},
	)
}
