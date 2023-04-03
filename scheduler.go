package main

import (
	"fmt"
	"os/exec"
	"sync"
	"time"
)

var (
	once      sync.Once
	Scheduler *scheduler
)

// 单例模式
func GetScheduler() *scheduler {
	once.Do(func() {
		Scheduler = &scheduler{
			pq:    NewPriorityQueue(3),
			count: 0,
			executors: map[string]Executor{
				GO:  &GoExecutor{},
				GRT: &GRTExecutor{},
				UTS: &UTSExecutor{},
			},
		}
	})
	return Scheduler
}

type scheduler struct {
	count     int
	pq        *PriorityQueue
	executors map[string]Executor
}

func (s *scheduler) WatchNewCommit() {
	// 监听新的git commit提交
	go func() {
		for {
			time.Sleep(time.Second)
		}
	}()
}

func (s *scheduler) WatchNewVersion() {
	// 监听是否刷新新版本
	go func() {
		for {
			time.Sleep(time.Second)
		}
	}()
}

func (s *scheduler) NewTask(taskID string, priority int) {
	t := NewGoTask(taskID, priority)

	s.pq.Push(t)
	s.count++
	fmt.Println("total task count: ", s.count)
}

func (s *scheduler) GetTaskCount() int {
	return s.count
}

func (s *scheduler) Start() {
	s.WatchNewCommit()
	s.WatchNewVersion()

	for {
		select {
		case <-s.pq.Waiting():
			t := s.Produce()
			s.Consume(t)
		}
	}
}

func (s *scheduler) Produce() Task {
	return s.pq.Produce()
}

func (s *scheduler) Consume(t Task) {
	s.executors[t.GetExecutor()].ExecuteTask(t)
	s.count -= 1
}

func Bash(cmd string) (status string, ret string) {
	status = "success"
	c := exec.Command("/bin/bash", "-c", cmd)
	out, err := c.CombinedOutput()
	if err != nil {
		fmt.Printf("combined out:\n%s\n", string(out))
		status = "failed"
		// log.Fatalf("cmd.Run() failed with %s\n", err)
	}

	ret = string(out)
	return
}
