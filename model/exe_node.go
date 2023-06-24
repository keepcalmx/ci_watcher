package model

import (
	"fmt"
)

// Node 执行节点
type Node struct {
	UUID       string
	Name       string
	IPAddr     string
	RootPwd    string
	SystemType string
	SystemArch string
}

func NewNode(UUID, name, ipAddr, rootPwd, systemType, systemArch string) Node {
	return Node{
		UUID:       UUID,
		Name:       name,
		IPAddr:     ipAddr,
		RootPwd:    rootPwd,
		SystemType: systemType,
		SystemArch: systemArch,
	}
}

// UninstallGaussDB 卸载 GaussDB
func (nm *Node) UninstallGaussDB() {
	// TODO: Implement
}

// InstallGaussDB 安装 GaussDB
func (nm *Node) InstallGaussDB() {
	// TODO: Implement
}

// InstallGRT 安装GRT
func (nm *Node) InstallGRT() {
	// TODO: Implement
}

// FlushVersion 刷新版本
func (nm *Node) FlushVersion() {
	// TODO: Implement
	cmd := "cat /etc/redhat-release"
	fmt.Println(cmd)
}

type VirEnv interface {
	Prepare()
	RunCase()
	InstallGRT()
	FlushVersion()
	InstallGaussDB()
	UninstallGaussDB()
}

type SingleEnv struct {
	Node       *Node
	DeployType string
}

func NewSingleEnv(node *Node, deployType string) *SingleEnv {
	return &SingleEnv{
		Node:       node,
		DeployType: deployType,
	}
}

func (s *SingleEnv) Prepare() {
	// TODO: Implement
}

func (s *SingleEnv) RunCase(c *TestCase) {
	// TODO: Implement
}

func (s *SingleEnv) InstallGRT() {
	// TODO: Implement
}

func (s *SingleEnv) FlushVersion() {
	// TODO: Implement
}

func (s *SingleEnv) InstallGaussDB() {
	// TODO: Implement
}

func (s *SingleEnv) UninstallGaussDB() {
	// TODO: Implement
}

type ClusterEnv struct {
	PrimaryNode  *Node
	StandbyNodes []*Node
	DeployType   string
}

func NewClusterEnv(primaryNode *Node, standbyNodes []*Node, deployType string) *ClusterEnv {
	return &ClusterEnv{
		PrimaryNode:  primaryNode,
		StandbyNodes: standbyNodes,
		DeployType:   deployType,
	}
}

func (s *ClusterEnv) Prepare() {
	// TODO: Implement
}

func (s *ClusterEnv) RunCase(c *TestCase) {
	// TODO: Implement
}

func (s *ClusterEnv) InstallGRT() {
	// TODO: Implement
}

func (s *ClusterEnv) FlushVersion() {
	// TODO: Implement
}

func (s *ClusterEnv) InstallGaussDB() {
	// TODO: Implement
}

func (s *ClusterEnv) UninstallGaussDB() {
	// TODO: Implement
}

// NodeManager 执行节点管理器
type NodeManager struct {
	Capacity     int
	Count        int
	Nodes        []Node
	RunningNodes []Node
	InitedNodes  map[string][]*Node
}

func NewNodeManager(capacity int) *NodeManager {
	return &NodeManager{
		Capacity: capacity,
		Count:    0,
		Nodes:    make([]Node, 0, capacity),
	}
}

func (nm *NodeManager) AddNode(node Node) {
	nm.Nodes = append(nm.Nodes, node)
	nm.Count++
}

func (nm *NodeManager) RemoveNode(UUID string) {
	for i, n := range nm.Nodes {
		if n.UUID == UUID {
			nm.Nodes = append(nm.Nodes[:i], nm.Nodes[i+1:]...)
			nm.Count--
			break
		}
	}
}

// TryPop 尝试弹出 num 个节点
func (nm *NodeManager) TryPop(num int) []Node {
	// 不满足期望个数，返回 nil
	if nm.Count == 0 || num > nm.Count {
		return nil
	}
	// 满足期望个数，返回前 num 个节点
	nodes := nm.Nodes[:num]
	nm.Nodes = nm.Nodes[num:]
	nm.Count -= num
	return nodes
}

func (nm *NodeManager) HasAvailableNode(t string) bool {
	return false
}

func (nm *NodeManager) HasStandbyNode(t string) bool {
	return false
}
