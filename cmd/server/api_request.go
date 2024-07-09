package main

import (
	"context"
	"fmt"

	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
)

// beforeRequest prints the HTTP method, base URL and URL path for the current request before executing it.
func beforeRequest(c *resty.Client, r *resty.Request) error {
	log.WithField("baseURL", c.BaseURL).Debugf("%s %s", r.Method, r.URL)
	return nil
}

// get is a generic utility function for reducing boilerplate client code.
func (s *Server) get(ctx context.Context, dst any, path string, expectedStatus int) (err error) {
	defer func() {
		if err != nil {
			log.Error(err)
		}
	}()
	req := s.api.R().SetResult(dst)
	var resp *resty.Response
	resp, err = req.SetContext(ctx).Get(path)
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
