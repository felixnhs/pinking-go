package api

import (
	"errors"
	"net/http"
	"pinking-go/lib"
	"pinking-go/lib/dao"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/apis"
)

type UserApi struct {
	dao *dao.UserDao
}

func BindUsersApi(provider *lib.DaoProvider, e *echo.Echo) {
	api := UserApi{}

	grp := e.Group("/profile", apis.RequireRecordAuth("users"))
	grp.GET("", api.getCurrent)
}

func (a *UserApi) getCurrent(c echo.Context) error {

	info := apis.RequestInfo(c)
	if info.AuthRecord == nil {
		return apis.NewForbiddenError("", errors.New(""))
	}

	user, err := a.dao.FindById(info.AuthRecord.Id)
	if err != nil {
		return err
	}

	canAccess, err := a.dao.Dao.CanAccessRecord(&user.Record, info, user.Collection().ViewRule)
	if !canAccess {
		return apis.NewForbiddenError("", err)
	}

	return c.JSON(http.StatusOK, user)
}
