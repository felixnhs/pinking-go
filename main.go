package main

import (
	"log"
	"pinking-go/server/api"
	"pinking-go/server/store"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

func main() {
	app := pocketbase.New()

	app.OnServe().BindFunc(func(e *core.ServeEvent) error {
		stores := store.BuildAllStores(e)

		api.BindUsersApi(e, stores)
		api.BindPostsApi(e, stores)

		return e.Next()
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
