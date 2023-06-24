package websocket

import (
	"fmt"
	"log"
	"sync"

	uuid "github.com/satori/go.uuid"
)

var (
	once sync.Once
	hub  *Hub
)

type Hub struct {
	// scheduler client
	schdClts *sync.Map
	// 来自scheduler的请求
	schdReqs chan *Request
	// 来自api的请求
	apiReqs chan *Request
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// GetHub returns a singleton hub
func GetHub() *Hub {
	once.Do(func() {
		hub = &Hub{
			schdClts: &sync.Map{},
			schdReqs: make(chan *Request, 100),
			apiReqs:  make(chan *Request, 100),
		}
		log.Println("hub is initialized...")
	})
	return hub
}

func (h *Hub) Register(client *SchedulerClient) {
	id := uuid.NewV1().String()
	h.schdClts.Store(id, client)
	client.requests = h.schdReqs
}

func (h *Hub) Unregister(uuid string) {
	h.schdClts.Delete(uuid)
}

func (h *Hub) IsOnline(uuid string) bool {
	_, ok := h.schdClts.Load(uuid)
	return ok
}

func (h *Hub) Run() {
	log.Println("hub is now running...")

	for {
		select {
		// 处理scheduler的请求
		case req := <-h.schdReqs:
			fmt.Println(req)
		// 处理api的请求
		case req := <-h.apiReqs:
			fmt.Println(req)
		}
	}
}
