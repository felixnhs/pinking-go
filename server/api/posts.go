package api

import (
	"pinking-go/server/api/model"
	"pinking-go/server/store"

	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

type PostApi struct {
	store *store.PostStore
}

func BindPostsApi(se *core.ServeEvent) {

	api := &PostApi{
		store: store.BuildPostStore(se),
	}

	// Auth
	grp := se.Router.Group("/posts")
	grp.Bind(apis.RequireAuth())
	grp.POST("/new", api.createNewPost)
}

func (a *PostApi) createNewPost(e *core.RequestEvent) error {

	req := model.CreatePostRequest{}

	if err := e.BindBody(&req); err != nil {
		return apis.NewBadRequestError("error_request_body", err)
	}

	post, err := a.store.CreatePost(e.Auth, &req)
	if err != nil {
		return apis.NewInternalServerError("error_create_post", err)
	}

	return RecordResponse(e, post)
}
