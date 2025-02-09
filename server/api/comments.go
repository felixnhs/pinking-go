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

	take, skip, err := utils.GetPaginationHeaders(e)
	if err != nil {
		return e.InternalServerError("error_request_info", err)
	}

	comments, err := a.Store().GetThread(id, take, skip)
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
	id := e.Request.PathValue("id")

	req := model.CreateCommentModel{}

	if err := e.BindBody(&req); err != nil {
		return apis.NewBadRequestError("error_request_body", err)
	}

	comment, err := a.Store().AddReplyToComment(e.Auth, id, &req)
	if err != nil {
		return apis.NewInternalServerError("error_create_comment", err)
	}

	return utils.RecordResponse(e, comment)
}
