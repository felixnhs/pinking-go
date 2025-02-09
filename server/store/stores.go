package store

import "github.com/pocketbase/pocketbase/core"

type StoreCollection struct {
	Users        UserStore
	Posts        PostStore
	Images       ImageStore
	Interactions InteractionsStore
	Comments     CommentStore
}

func BuildAllStores(se *core.ServeEvent) *StoreCollection {
	col := StoreCollection{}

	BuildUserStore(se, &col)
	BuildPostStore(se, &col)
	BuildInteractionsStore(se, &col)
	BuildImageStore(se, &col)
	BuildCommentStore(se, &col)

	return &col
}
