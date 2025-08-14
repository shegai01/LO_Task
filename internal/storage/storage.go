package storage

import (
	"log/slog"
	"sync"
	"sync/atomic"
	"time"

	"github.com/shegai01/LO_task/internal/model"
)

type Storage struct {
	mu         sync.RWMutex
	currId     atomic.Int64
	Task       map[int64]*model.Task
	someLogger *slog.Logger
}

func NewStorage() *Storage {
	return &Storage{
		Task: make(map[int64]*model.Task)}
}

func (s *Storage) Create(title string) (*model.Task, error) {
	id := s.currId.Add(1)
	task := &model.Task{
		ID:       id,
		Title:    title,
		Status:   "new",
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	}

	s.mu.Lock()
	s.Task[id] = task
	s.mu.Unlock()

	return task, nil
}

func (s *Storage) Get(id int64) (*model.Task, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	task, ok := s.Task[id]
	if !ok {
		return nil, false
	}

	return task, true

}
func (s *Storage) List(status model.Status) ([]*model.Task, error) {
	var arrTask []*model.Task

	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, task := range s.Task {
		if status != "" && task.Status != status {
			continue
		}

		arrTask = append(arrTask, task)
	}

	return arrTask, nil
}
