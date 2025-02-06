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

	items := []map[string]any{}
	for i := 0; i < len(r); i++ {
		items = append(items, r[i].WithCustomData(inc).PublicExport())
	}

	return e.JSON(http.StatusOK, items)
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
