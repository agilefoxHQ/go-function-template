package handlers

import (
	"net/http"
)

func (h *handler) middleware2(handler http.Handler) http.Handler {
	// Call the underlying handler
	return http.HandlerFunc(
		func(writer http.ResponseWriter, request *http.Request) {
			// And handle things on your own once in a while
			handler.ServeHTTP(writer, request)
		},
	)
}
