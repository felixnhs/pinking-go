package utils

import (
	"strconv"

	"github.com/pocketbase/pocketbase/core"
)

func GetPaginationHeaders(e *core.RequestEvent) (int, int, error) {
	info, err := e.RequestInfo()
	if err != nil {
		return 0, 0, e.InternalServerError("error_request_info", err)
	}

	take := GetQueryInt64(info, "take", 10)
	skip := GetQueryInt64(info, "skip", 0)

	return take, skip, nil
}

func GetQueryInt64(info *core.RequestInfo, name string, def int) int {
	str := info.Query[name]
	if str == "" {
		return def
	}

	val, err := strconv.Atoi(str)
	if err != nil {
		return def
	}

	return val
}

func GetQueryBool(info *core.RequestInfo, name string, def bool) bool {
	str := info.Query[name]
	if str == "" {
		return def
	}

	val, err := strconv.ParseBool(str)
	if err != nil {
		return def
	}

	return val
}
