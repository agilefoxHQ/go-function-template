package handlers

import (
	"fmt"
	"net/http"
)

func (h *handler) doThatThing(w http.ResponseWriter, r *http.Request) {
	// go do that thing
	_, _ = w.Write([]byte(fmt.Sprintf("%v!!", "🕺")))
	return
}
