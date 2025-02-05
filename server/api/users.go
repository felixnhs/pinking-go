package api

import (
	"pinking-go/lib/api/model"
	"pinking-go/lib/store"

	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

type UserApi struct {
	store *store.UserStore
}

func BindUsersApi(se *core.ServeEvent) {

	api := &UserApi{
		store: store.BuildUserStore(se),
	}

	se.Router.POST("/users/register", api.registerNewUser)

	grp := se.Router.Group("/users")
	grp.Bind(apis.RequireAuth())
	grp.GET("/me", api.getCurrentUser)
	grp.PUT("", api.updateUser)
}

func (a *UserApi) registerNewUser(e *core.RequestEvent) error {

	data := struct {
		Email    string `json:"email" form:"email"`
		Password string `json:"password" form:"password"`
	}{}

	if err := e.BindBody(&data); err != nil {
		return apis.NewBadRequestError("", err)
	}

	user, err := a.store.CreateNew(data.Email, data.Password)
	if err != nil {
		return apis.NewInternalServerError("error_register_user", err)
	}

	return apis.RecordAuthResponse(e, user.Record, "email", nil)
}

func (a *UserApi) getCurrentUser(e *core.RequestEvent) error {
	return RecordResponse(e, e.Auth)
}

func (a *UserApi) updateUser(e *core.RequestEvent) error {

	body := model.User{}

	if err := e.BindBody(&body); err != nil {
		return apis.NewBadRequestError("error_binding_input", err)
	}

	if err := a.store.UpdateUser(e.Auth, &body); err != nil {
		return apis.NewInternalServerError("error_updating_user", err)
	}

	return RecordResponse(e, e.Auth)
}
