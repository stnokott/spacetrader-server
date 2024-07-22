package main

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
	"github.com/stnokott/spacetrader/internal/api"
	"go.uber.org/ratelimit"
)

func configureRestyClient(r *resty.Client, baseURL string, token string) {
	r.
		SetBaseURL(baseURL).
		SetAuthToken(token).
		SetHeaders(map[string]string{
			"Accept":     "application/json",
			"User-Agent": "github.com/stnokott/spacetraders",
		}).
		SetTimeout(5 * time.Second). // TODO: allow configuring from env
		SetLogger(log.StandardLogger()).
		SetRetryAfter(func(_ *resty.Client, _ *resty.Response) (time.Duration, error) {
			return 5 * time.Second, nil
		}).
		SetRetryCount(3).
		OnBeforeRequest(beforeRequest())
}

func newRateLimiter() ratelimit.Limiter {
	// we want to be a little generous, the actual rate limit is 2 rps.
	return ratelimit.New(1)
}

// beforeRequest prints the HTTP method, base URL and URL path for the current request before executing it.
func beforeRequest() func(*resty.Client, *resty.Request) error {
	rateLimiter := newRateLimiter()
	return func(c *resty.Client, r *resty.Request) error {
		_ = rateLimiter.Take()
		log.WithField("baseURL", c.BaseURL).Debugf("%s %s", r.Method, r.URL)
		return nil
	}
}

// get is a generic utility function for reducing boilerplate client code.
func (s *Server) get(ctx context.Context, dst any, path string, expectedStatus int) (err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("%s: %w", path, err)
		}
	}()
	req := s.api.R().SetResult(dst)
	var resp *resty.Response
	resp, err = req.SetContext(ctx).Get(path)
	if err != nil {
		return
	}
	err = expectStatus(resp, expectedStatus)
	return
}

// UnexpectedStatusCodeErr is used when the HTTP status code returned from a request is different
// from the expected status code.
//
// `Msg` is filled from the response body.
type UnexpectedStatusCodeErr struct {
	Expected, Actual int
	Msg              string
}

func (e UnexpectedStatusCodeErr) Error() string {
	return fmt.Sprintf("expected status %d, got %d (%s)", e.Expected, e.Actual, e.Msg)
}

// expectStatus compares the status code of a request with the expected code and returns an error if a mismatch is encountered.
// Otherwise, it returns nil.
func expectStatus(resp *resty.Response, expectedStatus int) error {
	if resp.StatusCode() != expectedStatus {
		return UnexpectedStatusCodeErr{
			Expected: expectedStatus,
			Actual:   resp.StatusCode(),
			Msg:      string(resp.Body()),
		}
	}
	return nil
}

type pageFunc func(page int) (urlPath string)

type paginatedResult[T any] struct {
	Data T
	Err  error
}

// getPaginatedAsync asynchronously queries a paginated endpoint, traversing
// its pages.
// We use a goroutine for endpoints which return huge amounts of data across many pages
// like the "/systems" endpoint. Storing all the responses in one singular slice would
// consume a lot of memory at once.
// Sending this data via channels is the preferred method.
//
// API responses of type T are sent to the returned data channel.
//
// The internal goroutine can be stopped via the stop channel.
// This should only be done when an external error occurs, e.g. during
// consumption of the data channel.
// When an error occurs from within getPaginatedAsync, i.e. (<-data).Err != nil, the goroutine will stop automatically,
// there is no need to send a stop signal in that case.
//
// Pass a function pageFn which assembles the URL path (without the base URL) depending on
// the current page.
func getPaginatedAsync[T any](
	ctx context.Context,
	s *Server,
	pageFn pageFunc,
) (data <-chan paginatedResult[T], stop chan<- struct{}) {
	dataChan := make(chan paginatedResult[T]) // unbuffered since the API is expected to be slower than the consumer
	stopChan := make(chan struct{}, 1)        // buffered so the caller doesn't block when sending
	data, stop = dataChan, stopChan

	go func() {
		defer close(stopChan)
		defer close(dataChan)

		// total expected number of items
		total := 1 // start with total > 0 to enter the first loop iteration
		// total number of items received
		n := 0
		for page := 1; n < total; page++ {
			urlPath := pageFn(page)

			result := new(struct {
				Data []T       `json:"data"`
				Meta *api.Meta `json:"meta"`
			})
			if err := s.get(ctx, result, urlPath, 200); err != nil {
				dataChan <- paginatedResult[T]{
					Err: err,
				}
				return
			}
			for _, item := range result.Data {
				dataChan <- paginatedResult[T]{
					Data: item,
				}
			}
			// update the expected total item number
			total = result.Meta.Total
			// update the actual received item count so far
			n += len(result.Data)
			log.Infof("queried %d/%d of type %s", n, total, reflect.TypeOf(*new(T)).Name())

			// check cancel conditions (stop channel and context)
			select {
			case <-stopChan:
				return
			case <-ctx.Done():
				return
			default:
				continue
			}
		}
	}()

	return
}
