package model

// TestCase 测试用例
type TestCase struct {
	Name  string
	Owner string
	Path  string
	Type  string
}

func NewTestCase(name, owner, path string) TestCase {
	return TestCase{
		Name:  name,
		Owner: owner,
		Path:  path,
	}
}

type TestCaseManager struct {
	Capacity int
	Count    int
	Cases    chan *TestCase
}

func NewTestCaseManager(capacity int) *TestCaseManager {
	return &TestCaseManager{
		Capacity: capacity,
		Count:    0,
		Cases:    make(chan *TestCase, capacity),
	}
}

func (tcm *TestCaseManager) AddCase(c *TestCase) {
	tcm.Cases <- c
	tcm.Count++
}

func (tcm *TestCaseManager) Pop() *TestCase {
	c := <-tcm.Cases
	return c
}
