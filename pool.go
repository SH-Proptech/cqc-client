package cqc

import (
	"context"
	"sync"
)

// Task represents a function with its arguments.
type Task func(ctx context.Context) error

// WorkerPool manages the execution of tasks and collects errors.
type WorkerPool struct {
	workerCount int
	taskCh      chan Task
	errorCh     chan error // Channel to collect errors
	wg          sync.WaitGroup
}

// NewWorkerPool creates a new WorkerPool.
func NewWorkerPool(workerCount int) *WorkerPool {
	return &WorkerPool{
		workerCount: workerCount,
		taskCh:      make(chan Task),
		errorCh:     make(chan error, 10), // Use a buffered channel to prevent blocking
	}
}

// Start initializes the worker pool and starts processing tasks.
func (wp *WorkerPool) Start(ctx context.Context) {
	for i := 0; i < wp.workerCount; i++ {
		wp.wg.Add(1)
		go func() {
			defer wp.wg.Done()
			for task := range wp.taskCh {
				if err := task(ctx); err != nil {
					wp.errorCh <- err // Send errors to the error channel
				}
			}
		}()
	}
}

// Stop waits for all workers to finish and closes the error channel.
func (wp *WorkerPool) Stop() {
	close(wp.taskCh)  // Close task channel to signal workers to finish
	wp.wg.Wait()      // Wait for all workers to complete
	close(wp.errorCh) // Close the error channel after all tasks are processed
}

// Submit adds a new task to the worker pool.
func (wp *WorkerPool) Submit(task Task) {
	wp.taskCh <- task
}

// CollectErrors collects all errors from the error channel.
func (wp *WorkerPool) CollectErrors() []error {
	var errors []error
	// Use a separate wait group to wait for the error collection
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		for err := range wp.errorCh {
			errors = append(errors, err)
		}
	}()

	wg.Wait() // Wait for the error collection goroutine to finish
	return errors
}

func (wp *WorkerPool) IterateErrors(callback func(error)) {
	errors := wp.CollectErrors()
	for _, err := range errors {
		callback(err)
	}
}
