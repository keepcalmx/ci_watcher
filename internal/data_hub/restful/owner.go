package restful

import (
	"ci_watcher/internal/orm"
	"context"
	"encoding/json"

	"github.com/kataras/iris/v12"
	"go.mongodb.org/mongo-driver/bson"
)

func RegisterOwner(app *iris.Application) {
	app.Get("/owners", func(c iris.Context) {
		cursor, err := orm.OwnerColl.Find(context.TODO(), bson.D{})
		if err != nil {
			panic(err)
		}

		var results []orm.OwnerInfo
		if err = cursor.All(context.TODO(), &results); err != nil {
			panic(err)
		}

		for _, result := range results {
			cursor.Decode(&result)
			_, err := json.MarshalIndent(result, "", "    ")
			if err != nil {
				panic(err)
			}
		}

		c.JSON(results)
	})
}
