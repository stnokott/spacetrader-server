package mocks

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/stnokott/spacetrader-server/internal/api"
	"github.com/vitorsalgado/mocha/v3"
	"github.com/vitorsalgado/mocha/v3/expect"
	"github.com/vitorsalgado/mocha/v3/reply"
)

var m *mocha.Mocha

// StartMockAPI starts the mocked API server and returns its URL.
func StartMockAPI() (string, error) {
	fmt.Println("starting mocked API server")

	m = mocha.NewBasic()

	mockServerStatus()
	mockSystems()
	mockShips()

	url := m.Start().URL
	fmt.Printf("mock server running on URL %s\n", url)
	return url, nil
}

// StopMockAPI stops the mocked API server.
func StopMockAPI() error {
	fmt.Println("closing mocked API server")
	return m.Close()
}

func mockReplyOk(resp any) *reply.StdReply {
	return reply.
		OK().
		Delay(50*time.Millisecond).
		BodyJSON(resp).
		Header("content-type", "application/json").
		Header("charset", "utf-8")
}

var mockedServerStatus = &api.Status{
	Status:    "Online",
	Version:   "v0.0.0",
	LastReset: api.Date{Time: time.Now().Add(-24 * time.Hour)},
	Resets: api.ServerResets{
		Frequency: "never",
		Next:      time.Now().AddDate(1, 0, 0),
	},
}
var mockedSystems = GenerateSystems(5, -500, 500)
var mockedShips = []*api.Ship{
	NewShipInSystem("ENTERPRISE", mockedSystems[rand.Intn(len(mockedSystems))]),
	NewShipInSystem("FALCON", mockedSystems[rand.Intn(len(mockedSystems))]),
	NewShipInSystem("DISCOVERY", mockedSystems[rand.Intn(len(mockedSystems))]),
}

func mockServerStatus() {
	m.AddMocks(
		mocha.
			Get(expect.URLPath("/")).
			Name("server-status").
			Reply(mockReplyOk(mockedServerStatus)),
	)
}

type paginatedResponse[T any] struct {
	Data []T      `json:"data"`
	Meta api.Meta `json:"meta"`
}

func mockSystems() {
	systemsResponse := &paginatedResponse[*api.System]{
		Data: mockedSystems,
		Meta: api.Meta{
			Total: len(mockedSystems),
			Page:  1,
			Limit: len(mockedSystems),
		},
	}
	m.AddMocks(
		mocha.
			Get(expect.URLPath("/systems")).
			Name("list-systems").
			Reply(mockReplyOk(systemsResponse)),
	)
}

func mockShips() {
	fleetResponse := &paginatedResponse[*api.Ship]{
		Data: mockedShips,
		Meta: api.Meta{
			Total: len(mockedShips),
			Page:  1,
			Limit: len(mockedShips),
		},
	}
	m.AddMocks(
		mocha.
			Get(expect.URLPath("/my/ships")).
			Name("list-ships").
			Reply(mockReplyOk(fleetResponse)),
	)
}
