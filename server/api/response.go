package api

import (
	"net/http"

	"github.com/pocketbase/pocketbase/core"
)

func RecordResponse(e *core.RequestEvent, r *core.Record) error {
	return e.JSON(http.StatusOK, r.WithCustomData(true).PublicExport())
}

func EmptyResponse(e *core.RequestEvent) error {
	return e.JSON(http.StatusOK, struct{}{})
}

func TokenResponse(e *core.RequestEvent, token *string) error {
	return e.JSON(http.StatusOK, map[string]string{"token": *token})
}
