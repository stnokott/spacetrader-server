package main

import (
	"context"
	"errors"
	"fmt"
	"iter"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
	"github.com/stnokott/spacetrader-server/internal/api"
	"go.uber.org/ratelimit"
)

// TODO: query /my/agent on startup to validate token

func configureRestyClient(r *resty.Client, baseURL string, token string) {
	r.
		SetBaseURL(baseURL).
		SetAuthToken(token).
		SetHeaders(map[string]string{
			"Accept":     "application/json",
			"User-Agent": "github.com/stnokott/spacetrader-server",
		}).
		SetTimeout(5 * time.Second). // TODO: allow configuring from env
		SetLogger(log.StandardLogger()).
		SetRetryAfter(retryAfter).
		SetRetryCount(3).
		AddRetryCondition(func(r *resty.Response, _ error) bool {
			return r.StatusCode() == http.StatusTooManyRequests
		}).
		OnBeforeRequest(beforeRequest())
}

// retryAfter handles 429 rate-limiting responses and configures the wait time accordingly.
func retryAfter(_ *resty.Client, resp *resty.Response) (time.Duration, error) {
	if resp.StatusCode() == 429 {
		rateReset := resp.Header().Get("x-ratelimit-reset")
		if rateReset == "" {
			return 0, errors.New("got HTTP 429 without x-ratelimit-reset header")
		}
		t, err := time.Parse(time.RFC3339, rateReset)
		if err != nil {
			return 0, fmt.Errorf("parsing x-ratelimit-reset value: %w", err)
		}
		wait := t.Sub(time.Now().UTC())
		log.WithField("wait", wait).Debug("ratelimit exceeded")
		return wait, nil
	}
	return 5 * time.Second, nil
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

func newRateLimiter() ratelimit.Limiter {
	// 2 requests per second, according to https://docs.spacetraders.io/api-guide/rate-limits
	return ratelimit.New(2, ratelimit.Per(1*time.Second))
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

// getPaginated returns an iterator, traversing all pages of a paginated endpoint.
// Each iteration yields one page of data.
//
// Pass a function pageFn which assembles the URL path (without the base URL) depending on
// the current page.
//
// You should stop iterating when an error is yielded.
func getPaginated[T any](
	ctx context.Context,
	s *Server,
	pageFn pageFunc,
) iter.Seq2[[]T, error] {
	// total expected number of items
	total := 1 // start with total > 0 to enter the first loop iteration
	// total number of items received
	n := 0

	return func(yield func([]T, error) bool) {
		for page := 1; n < total; page++ {
			urlPath := pageFn(page)

			result := new(struct {
				Data []T       `json:"data"`
				Meta *api.Meta `json:"meta"`
			})
			if err := s.get(ctx, result, urlPath, 200); err != nil {
				_ = yield(nil, err)
				return
			}

			// update the expected total item number
			total = result.Meta.Total

			if !yield(result.Data, nil) {
				return
			}
			n += len(result.Data)
			log.Debugf("queried %03.0f%% (%d/%d) of type %T", float64(n)/float64(total)*100, n, total, *new(T))
		}
	}
}

// collectPages combines all pages of an iterator into one slice, returning any error encountered on any page.
func collectPages[V any](seq iter.Seq2[[]V, error]) ([]V, error) {
	out := []V{}
	for page, err := range seq {
		if err != nil {
			return nil, err
		}
		for _, item := range page {
			out = append(out, item)
		}
	}
	return out, nil
}
