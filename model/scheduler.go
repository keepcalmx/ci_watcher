package model

import (
	"fmt"
)

// Scheduler 调度节点执行测试用例
type Scheduler struct {
	// 状态 0：空闲 1：运行
	Status          int
	SingleEnvQueue  chan *SingleEnv
	ClusterEnvQueue chan *ClusterEnv
	// 空闲节点栈，用于重组执行环境
	NodeStack []*Node
	SQLCases  map[string][]*TestCase
	CLSCases  map[string][]*TestCase
}

func NewScheduler() *Scheduler {
	return &Scheduler{}
}

func (s *Scheduler) InitTestCase() {
	// TODO: Implement
}

func (s *Scheduler) InitExeEnv() {
	// 按照CENT_SQL、DIST_SQL、CENT_SSH、DIST_SSH的顺序轮流初始化环境
	// TODO: Implement
}

func (s *Scheduler) Init() {
	s.InitTestCase()
	s.InitExeEnv()
}

func (s *Scheduler) StartAll() {
	// TODO: Implement
	// 解析全部测试用例
	// 解析全部执行节点
	// 为每个执行节点分配测试用例
	// 如果有节点空余出来，继续分配测试用例

	for {
		select {
		case single := <-s.SingleEnvQueue:
			cases := s.SQLCases[single.DeployType]
			if len(cases) > 0 {
				nextCase := cases[0]
				single.RunCase(nextCase)
			} else {
				// 该组网的用例已经执行完毕，切换组网环境，优先在同类环境中切换
				// 例如：cent_ssh -> dist_ssh，dist_sql -> cent_sql

			}
		case cluster := <-s.ClusterEnvQueue:
			cases := s.CLSCases[cluster.DeployType]
			if len(cases) > 0 {
				nextCase := cases[0]
				cluster.RunCase(nextCase)
			} else {
				// 回收环境
			}
		default:
			fmt.Println("default")
		}
	}
}

func (s *Scheduler) StartOne() {
	// TODO: Implement
}

func (s *Scheduler) Pause() {
	// TODO: Implement
}

func (s *Scheduler) AddTestCase(testCase TestCase) {
	// s.TestCaseManager.AddTestCase(testCase)
}
