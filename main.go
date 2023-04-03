package main

import "ci_watcher/orm"

func main() {
	go orm.Migrate()

	go NewRESTfulAPI().Serve()

	GetScheduler().Start()
}
