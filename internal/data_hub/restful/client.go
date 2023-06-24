package restful

import (
	"ci_watcher/internal/data_hub/websocket"

	"github.com/kataras/iris/v12"
)

func RegisterClient(app *iris.Application) {
	app.Get("client", websocket.HandleNewSchedulerConn)
}
