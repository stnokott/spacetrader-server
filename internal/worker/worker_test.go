package worker

import (
	"context"
	"errors"
	"os"
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

func makeJob(progress float64, err error) (job Job, finish func()) {
	finishChan := make(chan struct{}, 1)
	job = func(ctx context.Context, progressChan chan<- float64) error {
		progressChan <- progress
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-finishChan:
			return err
		}
	}
	finish = func() {
		finishChan <- struct{}{}
	}
	return job, finish
}

func assertNoValue(assert *require.Assertions, ch <-chan error) {
	var hasValue bool
	select {
	case <-ch:
		hasValue = true
	default:
		hasValue = false
	}
	assert.False(hasValue)
}

func TestMain(m *testing.M) {
	log.SetLevel(log.DebugLevel)
	os.Exit(m.Run())
}

func TestErrors(t *testing.T) {
	assert := require.New(t)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// add job with nil error
	jobNil, finishNil := makeJob(0, nil)
	result := Add(ctx, "foo", jobNil)
	finishNil()
	assert.NoError(<-result)

	// add job with non-nil error
	jobErr, finishErr := makeJob(0, errors.New("bar"))
	result = Add(ctx, "bar", jobErr)
	finishErr()
	assert.Error(<-result)
}

func TestSequential(t *testing.T) {
	assert := require.New(t)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// add job
	firstJob, finishFirst := makeJob(0, nil)
	result := Add(ctx, "first", firstJob)
	// should have no result yet
	assertNoValue(assert, result)
	finishFirst()
	assert.NoError(<-result)

	// add second job
	secondJob, finishSecond := makeJob(0, errors.New("second"))
	result = Add(ctx, "second", secondJob)
	assertNoValue(assert, result)
	finishSecond()
	assert.Error(<-result)
}

func TestParallel(t *testing.T) {
	assert := require.New(t)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// add job
	firstJob, finishFirst := makeJob(0, nil)
	resultFirst := Add(ctx, "first", firstJob)

	// add second job
	secondJob, finishSecond := makeJob(0, nil)
	resultSecond := Add(ctx, "second", secondJob)

	assertNoValue(assert, resultFirst)
	assertNoValue(assert, resultSecond)

	finishSecond()
	assert.NoError(<-resultSecond)
	// first job should still be running
	assertNoValue(assert, resultFirst)

	finishFirst()
	assert.NoError(<-resultFirst)
}
