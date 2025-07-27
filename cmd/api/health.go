package main

import (
	"net/http"
)

// healthcheckHandler godoc
//
//	@Summary		Healthcheck
//	@Description	Healthcheck endpoint
//	@Tags			ops
//	@Produce		json
//	@Success		200	{object}	string	"ok"
//	@Router			/health [get]
func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// sending json response
	// w.Header().Set("Content-Type", "application/json")
	// w.Write([]byte(`{"status": "ok"}`))

	data := map[string]string{
		"status":  "ok",
		"env":     app.config.env,
		"version": version,
	}

	if err := writeJSON(w, http.StatusOK, data); err != nil {
		app.internalServerError(w, r, err)
		// writeJSONError(w, http.StatusInternalServerError, "err.Error()")
	}

	// needs fix
	// app.store.Posts.Create(r.Context())
}
