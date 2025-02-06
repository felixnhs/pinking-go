package api

import (
	"fmt"
	"pinking-go/server/api/model"
	"pinking-go/server/store"
	"strconv"

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
	grp.GET("", api.getPaginated)
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

func (a *PostApi) getPaginated(e *core.RequestEvent) error {

	info, err := e.RequestInfo()
	if err != nil {
		return e.InternalServerError("error_request_info", err)
	}

	take := getQueryInt64(info, "take", 10)
	skip := getQueryInt64(info, "skip", 0)

	posts, err := a.store.GetPosts(take, skip)
	if err != nil {
		return e.InternalServerError("error_retrieve_posts", err)
	}

	extended := getQueryBool(info, "extended", true)

	fmt.Printf("%v %v %v\n%v\n", take, skip, extended, posts)

	return MultipleRecordResponse(e, posts, extended)
}

func getQueryInt64(info *core.RequestInfo, name string, def int) int {
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

func getQueryBool(info *core.RequestInfo, name string, def bool) bool {
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
