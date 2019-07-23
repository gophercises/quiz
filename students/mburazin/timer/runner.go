package timer

import (
	"errors"
	"time"
)

// ErrorTimeout is a constant containing a timeout error
var ErrorTimeout = errors.New("Received timeout")

// Runner type is used to measure time for specified tasks.
// It will report timeout if given tasks don't finish in specified time.
// Use NewRunner() method to create a new Runner.
type Runner struct {
	timeout  <-chan time.Time
	complete chan error
	tasks    []func() error
}

// NewRunner creates a new runner of tasks that have to run before
// timeout given by d occurs.
func NewRunner(d time.Duration) *Runner {
	r := Runner{
		timeout:  time.After(d),
		complete: make(chan error),
	}

	return &r
}

// Add adds tasks which will be timed using the timer
func (r *Runner) Add(tasks ...func() error) {
	for _, task := range tasks {
		r.tasks = append(r.tasks, task)
	}
}

// Run runs the tasks previously added by Add() and starts the timer.
// If the tasks don't finish running in the specified duration given with the
// NewRunner() method, ErrorTimeout is returned.
// Otherwise nil is returned if all the tasks were successful, otherwise error
// describing the reason for task failure is returned.
func (r *Runner) Run() error {

	// run tasks
	go func() {
		r.complete <- r.run()
	}()

	// wait for tasks to finish, or for timeout
	select {
	case err := <-r.complete:
		return err
	case <-r.timeout:
		return ErrorTimeout
	}
}

func (r *Runner) run() error {
	for _, task := range r.tasks {
		err := task()
		if err != nil {
			return err
		}
	}

	return nil
}
