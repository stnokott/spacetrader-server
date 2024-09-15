// Package worker performs heavy-lifting tasks in the background.
package worker

import (
	"context"
	"fmt"
	"time"

	"github.com/stnokott/spacetrader-server/internal/log"
	"github.com/stnokott/spacetrader-server/internal/syncx"
	"go.uber.org/atomic"
)

var logger = log.ForComponent("worker")

// Worker maintains the job queue and processes jobs.
type Worker struct {
	// queue is a FIFO list of jobs, implemented by a simple channel.
	queue chan queueItem
	// index keeps track of the job names in the queue and maps them to their respective result channel.
	// Jobs are kept here until they have finished processing.
	index *syncx.Map[string, chan error]

	currentJob      atomic.String
	currentProgress atomic.Float64
}

type queueItem struct {
	Name string
	Job  Job
}

// Job performs a task.
type Job func(ctx context.Context, progressChan chan<- float64) error

// NewWorker returns a new job manager.
func NewWorker() *Worker {
	return &Worker{
		queue: make(chan queueItem),
		index: syncx.NewMap[string, chan error](),
	}
}

// Add adds a job to the end of the queue (FIFO).
//
// The returned channel can be queried to get the job result.
//
// A job name may only exist once in the queue.
// If a job with the given name already exists in the queue, a non-nil error will be returned and the queue will stay untouched.
func (w *Worker) Add(name string, job Job) (<-chan error, error) {
	_, exists := w.index.Get(name)
	if exists {
		return nil, fmt.Errorf("job with name '%s' currently processing or in the queue", name)
	}

	resultChan := make(chan error, 1) // buffered so jobs can finish without a consumer
	w.index.Set(name, resultChan)

	go func() {
		w.queue <- queueItem{
			Name: name,
			Job:  job,
		}
	}()

	return resultChan, nil
}

// AddAndWait adds the job to the queue and blocks until the job is finished.
//
// It returns the error from Add, if any, and otherwise, the error from the job.
func (w *Worker) AddAndWait(ctx context.Context, name string, job Job) error {
	result, err := w.Add(name, job)
	if err != nil {
		return err
	}
	select {
	case <-ctx.Done():
		return ctx.Err()
	case res := <-result:
		return res
	}
}

// Start consumes the queued jobs indefinitely.
// It is non-blocking.
//
// Consumption of jobs will stop when ctx expires.
//
// The worker will print the status of the current job in regular intervals if logState is true.
func (w *Worker) Start(ctx context.Context, logState bool) {
	go func() {
		// progressChan is used to communicate the progress of the current job.
		// it will be recreated for each new job to prevent deadlocks.
		var progressChan chan float64
		// doneChan is used internally to signal that the current job has finished.
		var doneChan chan struct{}

		if logState {
			go w.logState(ctx)
		}

		for {
			select {
			case <-ctx.Done():
				// abort when context expires
				close(w.queue)
				return
			case item := <-w.queue:
				progressChan = make(chan float64, 1) // buffered so we don't block jobs during consumption
				doneChan = make(chan struct{})       // unbuffered since only used internally
				go w.process(ctx, item, progressChan, doneChan)
			case prog := <-progressChan:
				w.currentProgress.Store(prog)
			case <-doneChan:
				// reset job status info
				w.currentJob.Store("")
				w.currentProgress.Store(0)
			}
		}
	}()
}

const logStateInterval = 10 * time.Second

func (w *Worker) logState(ctx context.Context) {
	ticker := time.NewTicker(logStateInterval)
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			currentJob := w.currentJob.Load()
			if currentJob != "" {
				logger.Infof("<%s> progress: %0.f%%", currentJob, w.currentProgress.Load()*100)
			} else {
				logger.Debug("no jobs")
			}
		}
	}
}

func (w *Worker) process(ctx context.Context, item queueItem, progressChan chan<- float64, doneChan chan<- struct{}) {
	w.currentJob.Store(item.Name)

	logger.Infof("<%s> started", item.Name)
	err := item.Job(ctx, progressChan)
	if err != nil {
		logger.Warnf("<%s> finished with err %v", item.Name, err)
	} else {
		logger.Infof("<%s> finished", item.Name)
	}

	// only clear the item from the index when processing is done (deferred),
	// to prevent having the same job type once in the queue and once in processing
	resultChan, _ := w.index.Pop(item.Name)
	resultChan <- err
	doneChan <- struct{}{}
	close(progressChan)
	close(resultChan)
	close(doneChan)
	return
}

// Status returns the current queue status.
func (w *Worker) Status() Status {
	return Status{
		CurrentJob:      w.currentJob.Load(),
		CurrentProgress: w.currentProgress.Load(),
	}
}

// Status represents the current queue status.
type Status struct {
	// CurrentJob is the empty string if there is no job active at the moment.
	CurrentJob string
	// CurrentProgress is the current job's progress (0-1).
	CurrentProgress float64
}
