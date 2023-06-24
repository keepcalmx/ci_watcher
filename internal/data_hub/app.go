package data_hub

import (
	"ci_watcher/internal/data_hub/restful"
	"log"

	"github.com/kataras/iris/v12"
)

type RESTfulAPI struct {
	app *iris.Application
}

func NewRESTfulAPI() *RESTfulAPI {
	return &RESTfulAPI{
		app: iris.New(),
	}
}

func (r *RESTfulAPI) RegisterAll() {
	restful.RegisterPing(r.app)
	restful.RegisterOwner(r.app)
	restful.RegisterCase(r.app)
	restful.RegisterTask(r.app)
	restful.RegisterClient(r.app)
}

func (r *RESTfulAPI) Serve() {
	log.Fatal(r.app.Listen(":9527"))
}
