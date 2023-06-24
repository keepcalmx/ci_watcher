package scheduler

import (
	"context"
	"fmt"
	"net"
	"os/exec"
	"sync"
	"time"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

var (
	once      sync.Once
	scheduler *Scheduler
)

// 单例模式
func GetScheduler() *Scheduler {
	once.Do(func() {
		conn, _, _, err := ws.DefaultDialer.Dial(
			context.TODO(),
			"ws://localhost:9527/client",
		)
		if err != nil {
			// handle error
		}

		scheduler = &Scheduler{
			pq:    NewPriorityQueue(3),
			count: 0,
			executors: map[string]Executor{
				GO:  &GoExecutor{},
				GRT: &GRTExecutor{},
				UTS: &UTSExecutor{},
			},
			wsConn: &conn,
		}
	})
	return scheduler
}

type Scheduler struct {
	count     int
	pq        *PriorityQueue
	executors map[string]Executor
	wsConn    *net.Conn
}

func (s *Scheduler) WatchNewCommit() {
	// 监听新的git commit提交
	for {
		time.Sleep(time.Second)
	}
}

func (s *Scheduler) WatchNewVersion() {
	// 监听是否刷新新版本
	for {
		time.Sleep(time.Second)
	}
}

// WatchMsgFromHub 监听来自DataHub的消息
func (s *Scheduler) WatchMsgFromHub() {
	for {
		msg, _, err := wsutil.ReadServerData(*s.wsConn)
		if err != nil {
			// handle error
		}

		// handle msg
		fmt.Print(msg)
	}
}

func (s *Scheduler) SendMsgToHub(msg []byte) {
	err := wsutil.WriteClientMessage(*s.wsConn, ws.OpText, msg)
	if err != nil {
		// handle error
	}
}

func (s *Scheduler) NewTask(taskID string, priority int) {
	t := NewGoTask(taskID, priority)

	s.pq.Push(t)
	s.count++
	fmt.Println("total task count: ", s.count)
}

func (s *Scheduler) GetTaskCount() int {
	return s.count
}

func (s *Scheduler) Start() {
	go s.WatchNewCommit()
	go s.WatchNewVersion()
	go s.WatchMsgFromHub()

	for range s.pq.Waiting() {
		t := s.Produce()
		s.Consume(t)
	}
	// for {
	// 	select {
	// 	case <-s.pq.Waiting():
	// 		t := s.Produce()
	// 		s.Consume(t)
	// 	}
	// }
}

func (s *Scheduler) Produce() Task {
	return s.pq.Produce()
}

func (s *Scheduler) Consume(t Task) {
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
