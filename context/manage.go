package context

import (
	req "og/reqeuest"
	"sync"
)

type Manager struct {
	Queue []*req.Request
	mu    sync.Mutex
}

func New() *Manager {
	queue := make([]*req.Request, 0)
	return &Manager{
		Queue: queue,
	}
}

func (self *Manager) Push(req *req.Request) {
	self.mu.Lock()
	defer self.mu.Unlock()
	self.Queue = append(self.Queue, req)
}

func (self *Manager) Pop() *req.Request {
	self.mu.Lock()
	defer self.mu.Unlock()
	r := self.Queue[0]
	self.Queue = self.Queue[1:]
	return r
}

func (self *Manager) Len() int {
	self.mu.Lock()
	defer self.mu.Unlock()
	return len(self.Queue)
}

func (self *Manager) Free() {
	self.mu.Lock()
	defer self.mu.Unlock()
	self.Queue = self.Queue[:0]
}
