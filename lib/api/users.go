package api

import (
	"errors"
	"net/http"
	"pinking-go/lib/dao"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

type UserApi struct {
	app *core.App
	dao *dao.UserDao
}

func BindUsersApi(app *core.App, e *echo.Echo) {
	api := UserApi{
		app: app,
		dao: &dao.UserDao{
			Dao: (*app).Dao(),
		},
	}

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

	canAccess, err := (*a.app).Dao().CanAccessRecord(&user.Record, info, user.Collection().ViewRule)
	if !canAccess {
		return apis.NewForbiddenError("", err)
	}

	return c.JSON(http.StatusOK, user)
}
