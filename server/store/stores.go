package store

import "github.com/pocketbase/pocketbase/core"

type StoreCollection struct {
	Users        *UserStore
	Posts        *PostStore
	Images       *ImageStore
	Interactions *InteractionsStore
}

func BuildAllStores(se *core.ServeEvent) *StoreCollection {
	u := BuildUserStore(se)
	inte := BuildInteractionsStore(se, u)
	img := BuildImageStore(se)
	posts := BuildPostStore(se, u, inte, img)

	return &StoreCollection{
		Users:        u,
		Posts:        posts,
		Images:       img,
		Interactions: inte,
	}
}
