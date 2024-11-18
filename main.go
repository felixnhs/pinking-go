package main

import (
	"log"
	"pinking-go/lib/api"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

func main() {
	app := pocketbase.New()

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		api.BindUsersApi(&app.App, e.Router)
		api.BindPostsApi(&app.App, e.Router)

		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
