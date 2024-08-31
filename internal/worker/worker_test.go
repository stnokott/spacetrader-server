package worker

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
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

func TestMain(m *testing.M) {
	log.SetLevel(log.DebugLevel)
	os.Exit(m.Run())
}

func TestWorker(t *testing.T) {
	assert := assert.New(t)

	m := NewWorker()

	m.Start(context.Background(), false)

	// add job with name A
	firstJob, finishFirst := makeJob(0, nil)
	firstResult, err := m.Add("first", firstJob)
	assert.NoError(err)

	// add job with name B, first job will be consumed from the queue so Add will block for a few ms max
	secondJob, finishSecond := makeJob(0, nil)
	secondResult, err := m.Add("second", secondJob)
	assert.NoError(err)

	// add job to full queue
	thirdJob, finishThird := makeJob(0, errors.New("we have encountered a foo"))
	go func() {
		// finish very first job
		finishFirst() // now, Add can return
	}()
	thirdResult, err := m.Add("third", thirdJob) // this will block until the first job has finished, freeing up one queue spot
	assert.NoError(err)

	finishSecond()
	finishThird()

	assert.NoError(<-firstResult)
	assert.NoError(<-secondResult)
	assert.Error(<-thirdResult)
}

func TestAdd(t *testing.T) {
	assert := assert.New(t)

	m := NewWorker()

	m.Start(context.Background(), false)

	// add job
	job, finish := makeJob(0, nil)
	_, err := m.Add("foo", job)
	assert.NoError(err)

	// add job with same key (not allowed while one with the same key is still running)
	_, err = m.Add("foo", job)
	assert.Error(err)

	// finish job
	finish()
	// now, we're allowed to add again
	time.Sleep(1 * time.Second) // give job time to finish
	_, err = m.Add("foo", job)
	assert.NoError(err)
}

func TestStatus(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	m := NewWorker()
	assertStatus := func(a assert.TestingT, expJob string, expProg float64) {
		status := m.Status()
		assert.Equal(a, expJob, status.CurrentJob)
		assert.Equal(a, expProg, status.CurrentProgress)
	}

	m.Start(ctx, false)
	assertStatus(t, "", 0)

	firstJob, finishFirst := makeJob(0.5, nil)

	m.Add("first", firstJob)
	assert.EventuallyWithT(t, func(c *assert.CollectT) {
		assertStatus(c, "first", 0.5)
	}, 1*time.Second, 100*time.Millisecond)

	finishFirst()
	assert.EventuallyWithT(t, func(c *assert.CollectT) {
		assertStatus(c, "", 0)
	}, 1*time.Second, 100*time.Millisecond)
}
