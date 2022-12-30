package handlers

import (
	"net/http"

	"github.com/rs/zerolog"

	"github.com/agilefoxHQ/go-function-template/store"
)

type handler struct {
	store  *store.Store
	logger *zerolog.Logger
}

func NewHandler(store *store.Store, logger *zerolog.Logger) *http.ServeMux {
	h := &handler{store: store, logger: logger}
	return handleRoutes(h)
}

func handleRoutes(h *handler) *http.ServeMux {
	mux := http.NewServeMux()
	handlerFunc := http.HandlerFunc(h.doThatThing)
	mux.Handle("/", h.middleware1(h.middleware2(handlerFunc)))
	return mux
}
