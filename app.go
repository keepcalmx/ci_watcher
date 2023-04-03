package main

import (
	"ci_watcher/orm"
	"context"
	"encoding/json"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"go.mongodb.org/mongo-driver/bson"
)

type RESTfulAPI struct {
	app       *fiber.App
	scheduler *scheduler
}

func NewRESTfulAPI() *RESTfulAPI {
	return &RESTfulAPI{
		app:       fiber.New(),
		scheduler: GetScheduler(),
	}
}

func (r *RESTfulAPI) Ping() {
	r.app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})
}

func (r *RESTfulAPI) Owner() {
	r.app.Get("/owners", func(c *fiber.Ctx) error {
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

		return c.JSON(results)
	})
}

func (r *RESTfulAPI) Case() {
	r.app.Get("/cases", func(c *fiber.Ctx) error {
		cursor, err := orm.CaseColl.Find(context.TODO(), bson.D{})
		if err != nil {
			panic(err)
		}

		var results []orm.CaseInfo
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

		return c.JSON(results)
	})
}

func (r *RESTfulAPI) Task() {
	r.app.Get("/tasks", func(c *fiber.Ctx) error {
		// 查询全部任务
		cases := []orm.CaseInfo{}
		cursor, err := orm.CaseColl.Find(context.TODO(), bson.D{})
		if err != nil {
			log.Fatal(err)
		}
		if err = cursor.All(context.TODO(), &cases); err != nil {
			log.Fatal(err)
		}

		return c.JSON(cases)
	})

	r.app.Post("/tasks", func(c *fiber.Ctx) error {
		task := struct {
			IDs []string `json:"ids"`
		}{}

		if err := c.BodyParser(&task); err != nil {
			return err
		}

		for _, id := range task.IDs {
			r.scheduler.NewTask(id, 1)
		}

		return c.JSON(task.IDs)
	})
}

func (r *RESTfulAPI) Metrics() {
	r.app.Get("/metrics", monitor.New(monitor.Config{
		Title: "MyService Metrics Page"},
	))

}

func (r *RESTfulAPI) Register() {
	r.Ping()
	r.Owner()
	r.Case()
	r.Task()
	r.Metrics()
}

func (r *RESTfulAPI) Serve() {
	r.Register()
	log.Fatal(r.app.Listen(":9000"))
}
