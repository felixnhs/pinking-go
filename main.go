package main

import (
	"log"
	"pinking-go/server/api"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

func main() {
	app := pocketbase.New()

	app.OnServe().BindFunc(func(e *core.ServeEvent) error {
		userApi := api.BindUsersApi(e)
		_ = api.BindPostsApi(e, userApi)

		return e.Next()
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
