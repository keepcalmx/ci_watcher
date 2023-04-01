package main

import (
	"fmt"
	"log"
	"os/exec"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

type Scheduler struct {
	Tasks chan Task
}

func NewScheduler() *Scheduler {
	return &Scheduler{
		Tasks: make(chan Task, 100),
	}
}

func (s *Scheduler) WatchNewCommit() {
	// 监听新的git commit提交

	go func() {
		for {
			time.Sleep(time.Second)
		}
	}()
}

func (s *Scheduler) WatchNewVersion() {
	// 监听是否刷新新版本
	go func() {
		for {
			time.Sleep(time.Second)
		}
	}()
}

func (s *Scheduler) WatchAPIs() {
	// 监听是否手动拉起任务
	go func() {

		app := fiber.New()

		app.Get("/ping", func(c *fiber.Ctx) error {
			return c.SendString("pong")
		})

		app.Get("/task", func(c *fiber.Ctx) error {
			s.Tasks <- Task{
				Name:    "go驱动支持pbe",
				WorkDir: "CENT/driver/go/roman",
				Command: "go test",
			}
			return c.JSON("success")
		})

		app.Get("/metrics", monitor.New(monitor.Config{Title: "MyService Metrics Page"}))

		log.Fatal(app.Listen(":9000"))
	}()
}

func (s *Scheduler) Start() {
	s.WatchNewCommit()
	s.WatchNewVersion()
	s.WatchAPIs()

	for {
		select {
		case t := <-s.Tasks:
			fmt.Println(1)
			// 调度用例执行
			fmt.Println(s.Bash("cd ./cases/" + t.WorkDir + " && go test"))
		}
	}

}

func (s *Scheduler) Bash(cmd string) string {
	c := exec.Command("/bin/bash", "-c", cmd)
	out, err := c.CombinedOutput()
	if err != nil {
		fmt.Printf("combined out:\n%s\n", string(out))
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	// fmt.Println(c.String(), ":")
	// fmt.Println(string(out))
	return string(out)
}

type Task struct {
	Name    string
	WorkDir string
	Command string
}
