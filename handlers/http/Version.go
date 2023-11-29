/*
 * Copyright © 2023 AgileFox GmbH,  Rizwan Iqbal <riz@agilefoxhq.com>
 */

package http

import (
	"net/http"

	"github.com/agilefoxHQ/go-function-template/config"
)

func (h *handler) Version(w http.ResponseWriter, r *http.Request) {
	WriteJSON(
		w, http.StatusOK, struct {
			Version string `json:"version"`
		}{
			Version: config.Version,
		},
	)
}
