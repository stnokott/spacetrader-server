// Package worker performs heavy-lifting tasks in the background.
package worker

import (
	"context"

	"github.com/stnokott/spacetrader-server/internal/log"
)

var logger = log.ForComponent("worker")

// Job performs a task.
type Job func(ctx context.Context, progressChan chan<- float64) error

// Add adds a job to the end of the queue (FIFO).
//
// The returned channel can be queried to get the job result.
func Add(ctx context.Context, name string, job Job) <-chan error {
	resultChan := make(chan error, 1)     // buffered so jobs can finish without a consumer
	progressChan := make(chan float64, 1) // buffered so we don't block jobs during consumption
	internalResultChan := make(chan error, 0)

	go func() {
		internalResultChan <- job(ctx, progressChan)
	}()

	go func() {
		logger.Infof("job <%s> started", name)
		var err error
		defer func() {
			if err != nil {
				logger.Warnf("job <%s> finished with err %v", name, err)
			} else {
				logger.Infof("job <%s> finished", name)
			}
			resultChan <- err
			close(resultChan)
			close(progressChan)
			close(internalResultChan)
		}()

		for {
			select {
			case <-ctx.Done():
				err = ctx.Err()
				return
			case result := <-internalResultChan:
				err = result
				return
			case p := <-progressChan:
				logger.Infof("job <%s> at %0.2f%%", name, p*100)
			}
		}
	}()

	return resultChan
}

// AddAndWait adds the job to the queue and blocks until the job is finished.
//
// It returns the error from the job or the context's error, whatever happens first.
func AddAndWait(ctx context.Context, name string, job Job) error {
	result := Add(ctx, name, job)

	select {
	case <-ctx.Done():
		return ctx.Err()
	case res := <-result:
		return res
	}
}
