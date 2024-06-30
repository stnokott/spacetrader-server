// Package client performs requests and unmarshals responses.
package client

import (
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
	"github.com/stnokott/spacetrader/internal/api"
)

// Client serves as an abstraction layer for the SpaceTraders API.
type Client struct {
	rest *resty.Client
}

// New creates and returns a new Client instance.
func New(baseURL string, token string) *Client {
	rest := resty.New()
	rest.SetBaseURL(baseURL)
	rest.SetAuthToken(token)
	rest.SetHeaders(map[string]string{
		"Accept":     "application/json",
		"User-Agent": "github.com/stnokott/spacetraders",
	})
	rest.SetTimeout(5 * time.Second) // TODO: allow configuring from env
	rest.SetLogger(log.StandardLogger())

	rest.OnBeforeRequest(beforeRequest)

	return &Client{
		rest: rest,
	}
}

// beforeRequest prints the HTTP method, base URL and URL path for the current request before executing it.
func beforeRequest(c *resty.Client, r *resty.Request) error {
	log.WithField("baseURL", c.BaseURL).Debugf("%s %s", r.Method, r.URL)
	return nil
}

// get is a generic utility function for reducing boilerplate client code.
func get[T any](rest *resty.Client, dst T, path string, expectedStatus int) error {
	req := rest.R().SetResult(dst)
	resp, err := req.Get(path)
	if err != nil {
		return err
	}
	if err = expectStatus(resp, expectedStatus); err != nil {
		return err
	}
	return nil
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
	return fmt.Sprintf("expected status %d, got %d", e.Expected, e.Actual)
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

// Status returns the current server status.
func (c *Client) Status() (*api.Status, error) {
	result := new(api.Status)
	if err := get(c.rest, result, "/", 200); err != nil {
		return nil, err
	}
	return result, nil
}
