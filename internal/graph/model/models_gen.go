// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"

	"github.com/nrfta/go-paging"
	"github.com/stnokott/spacetrader-server/internal/api"
)

type Agent struct {
	Name    string `json:"name"`
	Credits int64  `json:"credits"`
	Hq      string `json:"hq"`
}

type Jumpgate struct {
	From *Waypoint `json:"from"`
	To   *Waypoint `json:"to"`
	// waypoint ID for resolving waypoint
	FromWaypointID string `json:"-"`
	// waypoint ID for resolving waypoint
	ToWaypointID string `json:"-"`
}

type Query struct {
}

type Server struct {
	Version       string                `json:"version"`
	LastReset     time.Time             `json:"lastReset"`
	NextReset     time.Time             `json:"nextReset"`
	Stats         *ServerStats          `json:"stats"`
	Announcements []*ServerAnnouncement `json:"announcements"`
}

type ServerAnnouncement struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type ServerStats struct {
	Agents    int `json:"agents"`
	Ships     int `json:"ships"`
	Systems   int `json:"systems"`
	Waypoints int `json:"waypoints"`
}

type Ship struct {
	Name     string            `json:"name"`
	Role     api.ShipRole      `json:"role"`
	Status   api.ShipNavStatus `json:"status"`
	System   *System           `json:"system"`
	Waypoint *Waypoint         `json:"waypoint"`
	// system ID for resolving system
	SystemID string `json:"-"`
	// waypoint ID for resolving waypoint
	WaypointID string `json:"-"`
}

type System struct {
	Name         string              `json:"name"`
	Type         api.SystemType      `json:"type"`
	X            int                 `json:"x"`
	Y            int                 `json:"y"`
	Waypoints    []*Waypoint         `json:"waypoints"`
	Factions     []api.FactionSymbol `json:"factions"`
	HasJumpgates bool                `json:"hasJumpgates"`
}

type SystemConnection struct {
	Edges    []*SystemEdge    `json:"edges"`
	PageInfo *paging.PageInfo `json:"pageInfo"`
}

type SystemEdge struct {
	Cursor *string `json:"cursor,omitempty"`
	Node   *System `json:"node"`
}

type Waypoint struct {
	Name   string           `json:"name"`
	System *System          `json:"system"`
	Type   api.WaypointType `json:"type"`
	X      int              `json:"x"`
	Y      int              `json:"y"`
	// system ID for resolving system
	SystemID string `json:"-"`
}
