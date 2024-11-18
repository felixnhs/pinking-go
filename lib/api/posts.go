package api

import (
	"net/http"
	"pinking-go/lib/dao"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/core"
)

type PostApi struct {
	app *core.App
	dao *dao.PostDao
}

func BindPostsApi(app *core.App, e *echo.Echo) {
	api := PostApi{
		app: app,
		dao: &dao.PostDao{
			Dao: (*app).Dao(),
		},
	}

	grp := e.Group("/posts")
	grp.GET("", api.list)
	grp.GET("/:id", api.get)
}

func (a *PostApi) get(c echo.Context) error {
	id := c.PathParam("id")

	post, err := a.dao.FindById(id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, post)
}

func (a *PostApi) list(c echo.Context) error {
	posts, err := a.dao.FindLastPosts()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, posts)
}
