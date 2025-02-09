package store

import (
	"pinking-go/server/store/db"

	"github.com/pocketbase/pocketbase/core"
)

type InteractionsStore struct {
	app        *core.App
	collection *StoreCollection
}

func BuildInteractionsStore(se *core.ServeEvent, col *StoreCollection) {
	col.Interactions = InteractionsStore{
		app:        &se.App,
		collection: col,
	}
}

func (d *InteractionsStore) TableName() string {
	return "interactions"
}

func (d *InteractionsStore) AddLike(auth, post *core.Record) error {
	return d.addInteraction(auth, post, func(i *db.Interaction) { i.SetLikeType() })
}

func (d *InteractionsStore) AddUnlike(auth, post *core.Record) error {
	return d.addInteraction(auth, post, func(i *db.Interaction) { i.SetUnlikeType() })
}

func (d *InteractionsStore) AddComment(auth, post *core.Record) error {
	return d.addInteraction(auth, post, func(i *db.Interaction) { i.SetCommentType() })
}

func (d *InteractionsStore) AddShare(auth, post *core.Record) error {
	return d.addInteraction(auth, post, func(i *db.Interaction) { i.SetShareType() })
}

func (d *InteractionsStore) addInteraction(auth, post *core.Record, configure func(i *db.Interaction)) error {
	app := (*d.app)

	interationsCollection, err := app.FindCollectionByNameOrId(d.TableName())
	if err != nil {
		return err
	}

	inte := &db.Interaction{}
	inte.SetProxyRecord(core.NewRecord(interationsCollection))
	inte.SetPost(post.Id)
	inte.SetUser(auth.Id)

	configure(inte)

	return app.Save(inte)
}
