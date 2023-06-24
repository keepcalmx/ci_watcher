package restful

import (
	"ci_watcher/internal/orm"
	"context"
	"log"

	"github.com/kataras/iris/v12"
	"go.mongodb.org/mongo-driver/bson"
)

func RegisterTask(app *iris.Application) {
	app.Get("/tasks", func(c iris.Context) {
		// 查询全部任务
		cases := []orm.CaseInfo{}
		cursor, err := orm.CaseColl.Find(context.TODO(), bson.D{})
		if err != nil {
			log.Fatal(err)
		}
		if err = cursor.All(context.TODO(), &cases); err != nil {
			log.Fatal(err)
		}

		c.JSON(cases)
	})

	app.Post("/tasks", func(c iris.Context) {
		task := struct {
			IDs []string `json:"ids"`
		}{}

		if err := c.ReadBody(&task); err != nil {
			c.StopWithError(iris.StatusBadRequest, err)
			return
		}

		c.JSON(task.IDs)
	})
}
