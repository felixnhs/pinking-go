package api

import (
	"pinking-go/server/api/model"
	"pinking-go/server/store"
	"pinking-go/server/utils"

	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

type PostApi struct {
	stores *store.StoreCollection
}

func (a *PostApi) Store() *store.PostStore {
	return &a.stores.Posts
}

func BindPostsApi(se *core.ServeEvent, stores *store.StoreCollection) {

	api := &PostApi{
		stores: stores,
	}

	// Auth
	grp := se.Router.Group("/posts")
	grp.Bind(apis.RequireAuth(), RequireLockoutMiddleware())
	grp.POST("/new", api.createNewPost)
	grp.GET("", api.getPaginated)
	grp.POST("/{id}/like", api.likePost)
	grp.POST("/{id}/unlike", api.unlikePost)
	grp.GET("/{id}/comments", api.getComments)
}

func (a *PostApi) createNewPost(e *core.RequestEvent) error {

	req := model.CreatePostRequest{}

	if err := e.BindBody(&req); err != nil {
		return apis.NewBadRequestError("error_request_body", err)
	}

	post, err := a.Store().CreatePost(e.Auth, &req)
	if err != nil {
		return apis.NewInternalServerError("error_create_post", err)
	}

	return utils.RecordResponse(e, post)
}

func (a *PostApi) getPaginated(e *core.RequestEvent) error {

	take, skip, err := utils.GetPaginationHeaders(e)
	if err != nil {
		return e.InternalServerError("error_request_info", err)
	}

	posts, err := a.Store().GetPosts(e.Auth, take, skip)
	if err != nil {
		return e.InternalServerError("error_retrieve_posts", err)
	}

	return utils.MultipleRecordResponse(e, posts)
}

func (a *PostApi) likePost(e *core.RequestEvent) error {
	id := e.Request.PathValue("id")

	post, err := a.Store().LikePost(e.Auth, id)
	if err != nil {
		return e.InternalServerError("error_interact_post", err)
	}

	return utils.RecordResponse(e, post)
}

func (a *PostApi) unlikePost(e *core.RequestEvent) error {
	id := e.Request.PathValue("id")

	post, err := a.Store().UnlikePost(e.Auth, id)
	if err != nil {
		return e.InternalServerError("error_interact_post", err)
	}

	return utils.RecordResponse(e, post)
}

func (a *PostApi) getComments(e *core.RequestEvent) error {
	id := e.Request.PathValue("id")

	take, skip, err := utils.GetPaginationHeaders(e)
	if err != nil {
		return e.InternalServerError("error_request_info", err)
	}

	comments, err := a.stores.Comments.GetForPostPaginated(id, take, skip)
	if err != nil {
		return apis.NewInternalServerError("error_get_comments", err)
	}

	return utils.MultipleRecordResponse(e, comments)
}
