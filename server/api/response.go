package api

import (
	"net/http"

	"github.com/pocketbase/pocketbase/core"
)

func RecordResponse(e *core.RequestEvent, r *core.Record, includeCustom ...bool) error {
	inc := getIncludeCustomData(includeCustom...)
	return e.JSON(http.StatusOK, r.
		WithCustomData(inc).
		PublicExport())
}

func MultipleRecordResponse(e *core.RequestEvent, r []*core.Record, includeCustom ...bool) error {
	inc := getIncludeCustomData(includeCustom...)

	res := []map[string]any{}
	for _, rec := range r {
		res = append(res, rec.WithCustomData(inc).PublicExport())
	}

	return e.JSON(http.StatusOK, res)
}

func EmptyResponse(e *core.RequestEvent) error {
	return e.JSON(http.StatusOK, struct{}{})
}

func TokenResponse(e *core.RequestEvent, token *string) error {
	return e.JSON(http.StatusOK, map[string]string{"token": *token})
}

func getIncludeCustomData(includeCustom ...bool) bool {
	inc := true
	if len(includeCustom) > 0 {
		inc = includeCustom[0]
	}

	return inc
}
