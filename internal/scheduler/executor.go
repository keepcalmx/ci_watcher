package scheduler

import (
	"ci_watcher/internal/orm"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

type Executor interface {
	Install()
	ExecuteTask(Task)
	ReportResult()
	Uninstall()
}

// GoExecutor Go单测用例执行器
type GoExecutor struct {
	Version string
	result  string
	status  string
}

func (e *GoExecutor) Install() {
}

func (e *GoExecutor) ExecuteTask(t Task) {
	fmt.Println("GoTestTask Execute ", t.GetID())
	c := orm.CaseInfo{}
	// 1. 更新状态为running
	orm.CaseColl.FindOneAndUpdate(
		context.TODO(),
		bson.D{{Key: "id", Value: t.GetID()}},
		bson.D{{Key: "$set", Value: bson.D{
			{Key: "status", Value: "running"},
		},
		}},
	).Decode(&c)

	// 2. 执行go test
	status, ret := Bash("cd ./cases/" + c.WorkDir + " && go test")

	e.status, e.result = status, ret

	// 3. 更新状态为done
	orm.CaseColl.UpdateOne(
		context.TODO(),
		bson.D{{Key: "id", Value: t.GetID()}},
		bson.D{{Key: "$set", Value: bson.D{
			{Key: "status", Value: status},
			{Key: "result", Value: ret},
		},
		}},
	)
}

func (e *GoExecutor) ReportResult() {
	fmt.Println(e.status, e.result)
}

func (e *GoExecutor) Uninstall() {
}

// GRTExecutor GRT用例执行器
type GRTExecutor struct {
}

func (e *GRTExecutor) Install() {
}

func (e *GRTExecutor) ExecuteTask(t Task) {
}

func (e *GRTExecutor) ReportResult() {
}

func (e *GRTExecutor) Uninstall() {
}

// UTSExecutor UTS用例执行器
type UTSExecutor struct {
}

func (e *UTSExecutor) Install() {
}

func (e *UTSExecutor) ExecuteTask(t Task) {
}

func (e *UTSExecutor) ReportResult() {
}

func (e *UTSExecutor) Uninstall() {
}
