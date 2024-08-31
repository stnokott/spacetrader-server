package main

import (
	"net/http"
	"testing"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

func TestRetryAfter(t *testing.T) {
	respOk := &resty.Response{
		RawResponse: &http.Response{
			StatusCode: http.StatusOK,
		},
	}

	makeRespRateLimit := func() (*resty.Response, time.Duration) {
		d := 30 * time.Second
		header := http.Header{}
		header.Set("x-ratelimit-reset", time.Now().Add(d).Format(time.RFC3339))
		return &resty.Response{
			RawResponse: &http.Response{
				StatusCode: http.StatusTooManyRequests,
				Header:     header,
			},
		}, d
	}

	assert := assert.New(t)
	r := resty.New()

	// HTTP 200 should return default retry duration
	dRetry, err := retryAfter(r, respOk)
	assert.Nil(err)
	assert.Equal(5*time.Second, dRetry)

	// HTTP 429 should return a retry duration based on the header
	respRateLimit, dRl := makeRespRateLimit()
	dRetry, err = retryAfter(r, respRateLimit)
	dRetry += 1 * time.Second // account for test execution time
	assert.Nil(err)
	assert.GreaterOrEqual(dRetry, dRl)
}
