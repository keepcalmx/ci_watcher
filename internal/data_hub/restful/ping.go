package restful

import "github.com/kataras/iris/v12"

func RegisterPing(app *iris.Application) {
	// ping
	app.Get("/ping", func(c iris.Context) {
		c.WriteString("pong")
	})
}
