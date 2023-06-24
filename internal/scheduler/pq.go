package scheduler

import (
	"container/list"
	"fmt"
	"sync"
)

// PriorityQueue 优先级队列
type PriorityQueue struct {
	// 互斥锁
	mu sync.Mutex
	// 通知队列有新任务
	notice chan bool
	// 不同优先级的任务队列
	queues []*list.List
}

func NewPriorityQueue(MaxPriority int) *PriorityQueue {
	pq := &PriorityQueue{
		notice: make(chan bool, MEDIUMBUFFER),
		queues: make([]*list.List, 0, MaxPriority+1),
	}

	for i := 1; i <= MaxPriority+1; i++ {
		pq.queues = append(pq.queues, list.New())
	}

	return pq
}

// Push 新加入一个任务
func (pq *PriorityQueue) Push(t Task) error {
	pq.mu.Lock()
	defer pq.mu.Unlock()

	p := t.GetPriority()

	if p < 1 || p > len(pq.queues) {
		return fmt.Errorf("invalid priority: %d", p)
	}

	pq.queues[p].PushBack(t)
	pq.notice <- true
	return nil
}

// AddToTop 将任务添加到队列头部(下一个执行)
func (pq *PriorityQueue) AddToTop(t *Task) {
	pq.mu.Lock()
	defer pq.mu.Unlock()

	pq.queues[1].PushFront(t)
	pq.notice <- true
}

// Produce 从队列中取出下一个最高优先级的任务
func (pq *PriorityQueue) Produce() Task {
	pq.mu.Lock()
	defer pq.mu.Unlock()

	for _, queue := range pq.queues {
		if queue.Len() > 0 {
			e := queue.Front()
			queue.Remove(e)
			return e.Value.(Task)
		}
	}

	return nil
}

// Waiting 返回一个通道，当队列中有新任务时，通道会有通知
func (pq *PriorityQueue) Waiting() <-chan bool {
	return pq.notice
}
