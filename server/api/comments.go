package api

import (
	"pinking-go/server/api/model"
	"pinking-go/server/store"
	"pinking-go/server/utils"

	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

type CommentApi struct {
	stores *store.StoreCollection
}

func (a *CommentApi) Store() *store.CommentStore {
	return &a.stores.Comments
}

func BindCommentApi(se *core.ServeEvent, stores *store.StoreCollection) {

	api := &CommentApi{
		stores: stores,
	}

	// Auth
	grp := se.Router.Group("/comments")
	grp.Bind(apis.RequireAuth(), RequireLockoutMiddleware())
	grp.POST("/{id}/reply", api.postReply)
	grp.POST("/new", api.createComment)
	grp.GET("/{id}", api.getCommentThread)

}

func (a *CommentApi) getCommentThread(e *core.RequestEvent) error {
	id := e.Request.PathValue("id")

	info, err := e.RequestInfo()
	if err != nil {
		return e.InternalServerError("error_request_info", err)
	}

	take := getQueryInt64(info, "take", 10)
	skip := getQueryInt64(info, "skip", 0)

	comments, err := a.Store().GetThread(e.Auth, id, take, skip)
	if err != nil {
		return e.InternalServerError("error_retrieve_comments", err)
	}

	return utils.MultipleRecordResponse(e, comments)
}

func (a *CommentApi) createComment(e *core.RequestEvent) error {
	req := model.CreateCommentModel{}

	if err := e.BindBody(&req); err != nil {
		return apis.NewBadRequestError("error_request_body", err)
	}

	comment, err := a.Store().AddToPost(e.Auth, &req)
	if err != nil {
		return apis.NewInternalServerError("error_create_comment", err)
	}

	return utils.RecordResponse(e, comment)
}

func (a *CommentApi) postReply(e *core.RequestEvent) error {
	req := model.CreateReplyModel{}

	if err := e.BindBody(&req); err != nil {
		return apis.NewBadRequestError("error_request_body", err)
	}

	comment, err := a.Store().AddReplyToComment(e.Auth, &req)
	if err != nil {
		return apis.NewInternalServerError("error_create_comment", err)
	}

	return utils.RecordResponse(e, comment)
}
