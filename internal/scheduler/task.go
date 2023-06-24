package scheduler

import "ci_watcher/internal/orm"

// Task 任务接口
type Task interface {
	GetID() string
	GetPriority() int
	GetExecutor() string
}

// GoTask Go单测任务
type GoTask struct {
	orm.CaseInfo
	ID       string
	Err      error
	priority int
}

func NewGoTask(id string, priority int) *GoTask {
	return &GoTask{
		ID:       id,
		priority: priority,
	}
}

func (t *GoTask) GetID() string {
	return t.ID
}

func (t *GoTask) GetPriority() int {
	return t.priority
}

func (t *GoTask) GetExecutor() string {
	return GO
}
