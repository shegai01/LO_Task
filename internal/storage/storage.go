package storage

import (
	"example/internal/model"
	"sync"
	"sync/atomic"
	"time"
)

type Storage struct {
	mu   sync.RWMutex
	rand atomic.Int64
	Task map[int64]*model.Task
}

func NewStorage() *Storage {
	return &Storage{
		Task: make(map[int64]*model.Task)}
}
func (s *Storage) Create(title string) (*model.Task, error) {
	id := s.rand.Add(1)
	task := &model.Task{
		ID:       id,
		Title:    title,
		CreateAt: time.Now(),
	}
	s.mu.RLock()
	s.Task[id] = task
	s.mu.RUnlock()
	return task, nil
}
func (s *Storage) Get(id int64) (*model.Task, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	t, ok := s.Task[id]
	if ok {
		return t, true
	}
	return nil, false

}
