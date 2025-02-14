package api

import (
	"pinking-go/server/api/model"
	"pinking-go/server/store"
	"pinking-go/server/store/db"
	"pinking-go/server/utils"

	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

type UserApi struct {
	stores *store.StoreCollection
}

func (a *UserApi) Store() *store.UserStore {
	return &a.stores.Users
}

func BindUsersApi(se *core.ServeEvent, stores *store.StoreCollection) {

	api := &UserApi{
		stores: stores,
	}

	// Anonym
	se.Router.POST("/users/register", api.registerNewUser)
	se.Router.POST("/users/login", api.login)

	// Auth
	grp := se.Router.Group("/users")
	grp.Bind(apis.RequireAuth(), RequireLockoutMiddleware())
	grp.POST("/logout", api.logout)
	grp.POST("/resetpassword", api.resetPassword)
	grp.GET("/me", api.getCurrentUser)
	grp.GET("/me/followers", api.getCurrentUserFollowers)
	grp.GET("/me/following", api.getCurrentUserFollowings)
	grp.GET("/{id}", api.getProfile)
	grp.GET("/{id}/posts", api.getPosts)
	grp.PUT("", api.updateUser)
	grp.PUT("/me/avatar", api.updateAvatar)
	grp.DELETE("/me/avatar", api.deleteAvatar)
	grp.POST("/{id}/follow", api.followUser)
	grp.POST("/{id}/unfollow", api.unfollowUser)
}

func (a *UserApi) registerNewUser(e *core.RequestEvent) error {

	req := model.RegistrationLoginRequest{}

	if err := e.BindBody(&req); err != nil {
		return apis.NewBadRequestError("error_request_body", err)
	}

	user, err := a.Store().CreateNew(req.Email, req.Password)
	if err != nil {
		return apis.NewInternalServerError("error_register_user", err)
	}

	return apis.RecordAuthResponse(e, user, "email", nil)
}

func (a *UserApi) login(e *core.RequestEvent) error {

	req := model.RegistrationLoginRequest{}

	if err := e.BindBody(&req); err != nil {
		return apis.NewBadRequestError("error_request_body", err)
	}

	user, err := a.Store().FindByEmail(req.Email)
	if err != nil {
		return apis.NewForbiddenError("error_user_login", err)
	}

	if !user.ValidatePassword(req.Password) {
		return apis.NewForbiddenError("error_user_login", user.Id)
	}

	if user.GetLockoutEnabled() {
		return apis.NewForbiddenError("error_user_lockout", user.Id)
	}

	return apis.RecordAuthResponse(e, user.Record, "email", nil)
}

func (a *UserApi) logout(e *core.RequestEvent) error {

	e.Auth.RefreshTokenKey()

	if err := e.App.Save(e.Auth); err != nil {
		return apis.NewInternalServerError("error_db_action", err)
	}

	return utils.EmptyResponse(e)
}

func (a *UserApi) resetPassword(e *core.RequestEvent) error {

	req := model.ResetPasswordRequest{}

	if err := e.BindBody(&req); err != nil {
		return apis.NewBadRequestError("error_request_body", err)
	}

	token, err := a.Store().ResetPassword(e.Auth, req.OldPassword, req.NewPassword)
	if err != nil {
		return apis.NewForbiddenError("error_reset_password", err)
	}

	return utils.TokenResponse(e, token)
}

func (a *UserApi) getCurrentUser(e *core.RequestEvent) error {
	user, err := a.Store().GetById(e.Auth.Id)
	if err != nil {
		return apis.NewInternalServerError("error_get_current_auth", err)
	}

	user.Unhide(db.User_Bio)

	return utils.RecordResponse(e, user.Record)
}

func (a *UserApi) getProfile(e *core.RequestEvent) error {
	id := e.Request.PathValue("id")

	user, err := a.Store().GetById(id)
	if err != nil {
		return apis.NewInternalServerError("error_get_profile", err)
	}

	user.Unhide(db.User_Bio)

	return utils.RecordResponse(e, user.Record)
}

func (a *UserApi) updateUser(e *core.RequestEvent) error {

	body := model.UserRequest{}

	if err := e.BindBody(&body); err != nil {
		return apis.NewBadRequestError("error_request_body", err)
	}

	if err := a.Store().UpdateUser(e.Auth, &body); err != nil {
		return apis.NewInternalServerError("error_updating_user", err)
	}

	return utils.RecordResponse(e, e.Auth)
}

func (a *UserApi) updateAvatar(e *core.RequestEvent) error {

	body := model.AvatarRequest{}

	if err := e.BindBody(&body); err != nil {
		return apis.NewBadRequestError("error_request_body", err)
	}

	if err := a.Store().UpdateAvatar(e.Auth, body.Base64); err != nil {
		return apis.NewBadRequestError("error_update_avatar", err)
	}

	return utils.RecordResponse(e, e.Auth)
}

func (a *UserApi) deleteAvatar(e *core.RequestEvent) error {
	if err := a.Store().ClearAvatar(e.Auth); err != nil {
		return apis.NewBadRequestError("error_clear_avatar", err)
	}

	return utils.RecordResponse(e, e.Auth)
}

func (a *UserApi) getPosts(e *core.RequestEvent) error {
	id := e.Request.PathValue("id")

	take, skip, err := utils.GetPaginationHeaders(e)
	if err != nil {
		return e.InternalServerError("error_request_info", err)
	}

	posts, err := a.stores.Posts.GetPostsForUser(e.Auth, id, take, skip)
	if err != nil {
		return e.InternalServerError("error_retrieve_posts", err)
	}

	return utils.MultipleRecordResponse(e, posts)
}

func (a *UserApi) followUser(e *core.RequestEvent) error {

	id := e.Request.PathValue("id")

	if err := a.stores.Users.AddFollow(e.Auth, id); err != nil {
		return apis.NewInternalServerError("error_follow_user", err)
	}

	return utils.EmptyResponse(e)
}

func (a *UserApi) unfollowUser(e *core.RequestEvent) error {

	id := e.Request.PathValue("id")

	if err := a.stores.Users.RemoveFollow(e.Auth, id); err != nil {
		return apis.NewInternalServerError("error_unfollow_user", err)
	}

	return utils.EmptyResponse(e)
}

func (a *UserApi) getCurrentUserFollowers(e *core.RequestEvent) error {
	records, err := a.stores.Users.GetFollowers(e.Auth)
	if err != nil {
		return apis.NewInternalServerError("error_retrieve_followers", err)
	}

	return utils.MultipleRecordResponse(e, records)
}

func (a *UserApi) getCurrentUserFollowings(e *core.RequestEvent) error {
	records, err := a.stores.Users.GetFollowingUsers(e.Auth)
	if err != nil {
		return apis.NewInternalServerError("error_retrieve_following", err)
	}

	return utils.MultipleRecordResponse(e, records)
}
