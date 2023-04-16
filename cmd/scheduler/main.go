package main

import (
	"ci_watcher/internal/scheduler"
)

func main() {
	scheduler.GetScheduler().Start()
}
