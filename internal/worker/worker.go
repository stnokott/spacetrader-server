// Package worker performs heavy-lifting tasks in the background.
package worker

import (
	"context"
	"time"

	"github.com/stnokott/spacetrader-server/internal/log"
)

var logger = log.ForComponent("worker")

// Job performs a task.
type Job func(ctx context.Context, progressChan chan<- float64) error

// Add adds a job to the end of the queue (FIFO).
//
// The returned channel can be queried to get the job result.
func Add(ctx context.Context, name string, job Job, options ...Option) <-chan error {
	o := makeOptions(options)

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
				logger.Warnf("job <%s> finished with err: %v", name, err)
			} else {
				logger.Infof("job <%s> finished", name)
			}
			resultChan <- err
			close(resultChan)
			close(progressChan)
			close(internalResultChan)
		}()

		logTimer := time.NewTicker(o.maxLogFrequency)
		for {
			select {
			case <-ctx.Done():
				err = ctx.Err()
				return
			case result := <-internalResultChan:
				err = result
				return
			case p := <-progressChan:
				// thottle progress logging by checking for timer expiration
				select {
				case <-logTimer.C:
					logger.Infof("job <%s> at %0.2f%%", name, p*100)
				default:
				}
			}
		}
	}()

	return resultChan
}

// AddAndWait adds the job to the queue and blocks until the job is finished.
//
// It returns the error from the job or the context's error, whatever happens first.
func AddAndWait(ctx context.Context, name string, job Job, options ...Option) error {
	result := Add(ctx, name, job, options...)

	select {
	case <-ctx.Done():
		return ctx.Err()
	case res := <-result:
		return res
	}
}

type options struct {
	// maxLogFrequency limits logging of any job updates (i.e. progress) to only one call per duration. Default is 5s.
	maxLogFrequency time.Duration
}

type Option func(o *options)

func makeOptions(opts []Option) *options {
	// create options instance with default values
	opt := &options{
		maxLogFrequency: 5 * time.Second,
	}
	for _, o := range opts {
		o(opt)
	}
	return opt
}

func WithMaxLogFrequency(d time.Duration) Option {
	return func(o *options) {
		o.maxLogFrequency = d
	}
}
