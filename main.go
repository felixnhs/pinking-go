package main

import (
	"log"
	"pinking-go/lib/api"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

func main() {
	app := pocketbase.New()

	// app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
	// 	// api.BindUsersApi(customApi, e.Router)
	// 	// customApi.PostApi = api.BuildPostsApi(&app.App, e.Router)

	// 	provider := lib.BuildProvider(&app.App)

	// 	api.BindUsersApi(provider, e.Router)

	// 	return nil
	// })

	app.OnServe().BindFunc(func(e *core.ServeEvent) error {
		api.BindUsersApi(e)

		return e.Next()
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
