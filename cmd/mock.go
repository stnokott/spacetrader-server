//go:build mock

package main

import "github.com/stnokott/spacetrader-server/tests/mocks"

func init() {
	beforeHooks = append(beforeHooks, startMockAPI)
	afterHooks = append(afterHooks, stopMockAPI)
}

func startMockAPI() error {
	url, err := mocks.StartMockAPI()
	if err != nil {
		return err
	}
	baseURL = url
	return nil
}

func stopMockAPI() error {
	return mocks.StopMockAPI()
}
