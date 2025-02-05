package api

import (
	"net/http"

	"github.com/pocketbase/pocketbase/core"
)

func RecordResponse(e *core.RequestEvent, r *core.Record) error {
	return e.JSON(http.StatusOK, r.WithCustomData(true).PublicExport())
}
