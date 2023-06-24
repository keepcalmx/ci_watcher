package main

import (
	"ci_watcher/internal/data_hub"
	"ci_watcher/internal/orm"
)

func main() {
	go orm.Migrate()
	go data_hub.NewRESTfulAPI().Serve()
}
