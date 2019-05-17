// Copyright 2019 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package hippo

import (
	"sync"
)

// WorkersPool struct
type WorkersPool struct {
	Tasks       []*Task
	concurrency int
	tasksChan   chan *Task
	wg          sync.WaitGroup
}

// Task struct
type Task struct {
	Err    error
	Result string
	f      func() (string, error)
}

// NewWorkersPool initializes a new pool with the given tasks
func NewWorkersPool(tasks []*Task, concurrency int) *WorkersPool {
	return &WorkersPool{
		Tasks:       tasks,
		concurrency: concurrency,
		tasksChan:   make(chan *Task),
	}
}

// Run runs all work within the pool and blocks until it's finished.
func (w *WorkersPool) Run() {
	for i := 0; i < w.concurrency; i++ {
		go w.work()
	}

	w.wg.Add(len(w.Tasks))
	for _, task := range w.Tasks {
		w.tasksChan <- task
	}

	// all workers return
	close(w.tasksChan)

	w.wg.Wait()
}

// The work loop for any single goroutine.
func (w *WorkersPool) work() {
	for task := range w.tasksChan {
		task.Run(&w.wg)
	}
}

// NewTask initializes a new task based on a given work
func NewTask(f func() (string, error)) *Task {
	return &Task{f: f}
}

// Run runs a Task
func (t *Task) Run(wg *sync.WaitGroup) {
	t.Result, t.Err = t.f()
	wg.Done()
}
