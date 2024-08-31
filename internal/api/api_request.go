package api

import (
	"context"
	"errors"
	"fmt"
	"iter"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/stnokott/spacetrader-server/internal/log"
	"go.uber.org/ratelimit"
)

// TODO: query /my/agent on startup to validate token

var logger = log.ForComponent("api")

type Client struct {
	r *resty.Client
}

func NewClient(baseURL string, token string) *Client {
	r := resty.New()

	r.
		SetBaseURL(baseURL).
		SetAuthToken(token).
		SetHeaders(map[string]string{
			"Accept":     "application/json",
			"User-Agent": "github.com/stnokott/spacetrader-server",
		}).
		SetTimeout(10 * time.Second). // TODO: allow configuring from env
		SetLogger(logger).
		SetRetryAfter(retryAfter).
		SetRetryCount(5).
		AddRetryCondition(func(r *resty.Response, _ error) bool {
			return r.StatusCode() == http.StatusTooManyRequests
		}).
		OnBeforeRequest(beforeRequest())

	return &Client{r}
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
		logger.Debugf("ratelimit exceeded, waiting %s", wait)
		return wait, nil
	}
	return 5 * time.Second, nil
}

// beforeRequest prints the HTTP method, base URL and URL path for the current request before executing it.
func beforeRequest() func(*resty.Client, *resty.Request) error {
	rateLimiter := newRateLimiter()
	return func(c *resty.Client, r *resty.Request) error {
		_ = rateLimiter.Take()
		logger.Debugf("%s %s", r.Method, r.URL)
		return nil
	}
}

func newRateLimiter() ratelimit.Limiter {
	// 2 requests per second, according to https://docs.spacetraders.io/api-guide/rate-limits
	return ratelimit.New(2, ratelimit.Per(1*time.Second))
}

// Get is a generic utility function for reducing boilerplate client code.
func (c *Client) Get(ctx context.Context, dst any, path string) (err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("%s: %w", path, err)
		}
	}()
	req := c.r.R().SetResult(dst)
	var resp *resty.Response
	resp, err = req.SetContext(ctx).Get(path)
	if err != nil {
		return
	}
	if !resp.IsSuccess() {
		err = fmt.Errorf("unexpected status code %d", resp.StatusCode())
	}
	return
}

type pageFunc func(page int) (urlPath string)

type pageResult[T any] struct {
	Items []T
	// The expected total number of items returned once the iterator is exhausted.
	Total int
}

// GetPaginated returns an iterator, traversing all pages of a paginated endpoint.
// Each iteration yields one page of data.
//
// Pass a function pageFn which assembles the URL path (without the base URL) depending on
// the current page.
//
// The iterator will stop when an error is returned.
func GetPaginated[T any](
	ctx context.Context,
	c *Client,
	pageFn pageFunc,
) iter.Seq2[*pageResult[T], error] {
	// total expected number of items
	total := 1 // start with total > 0 to enter the first loop iteration
	// total number of items received
	n := 0

	return func(yield func(*pageResult[T], error) bool) {
		for page := 1; n < total; page++ {
			urlPath := pageFn(page)

			result := new(struct {
				Data []T   `json:"data"`
				Meta *Meta `json:"meta"`
			})
			if err := c.Get(ctx, result, urlPath); err != nil {
				_ = yield(nil, err)
				return
			}

			// update the expected total item number
			total = result.Meta.Total
			if total == 0 {
				return
			}

			if !yield(&pageResult[T]{
				Items: result.Data,
				Total: total,
			}, nil) {
				return
			}
			n += len(result.Data)
			logger.Debugf("queried %03.0f%% (%d/%d) of type %T", float64(n)/float64(total)*100, n, total, *new(T))
		}
	}
}

// CollectPages combines all pages of an iterator into one slice, returning any error encountered on any page.
func CollectPages[V any](seq iter.Seq2[*pageResult[V], error]) ([]V, error) {
	var out []V

	i := 0
	for page, err := range seq {
		if err != nil {
			return nil, err
		}
		if out == nil {
			out = make([]V, page.Total)
		}
		for _, item := range page.Items {
			out[i] = item
			i++
		}
	}
	return out, nil
}
