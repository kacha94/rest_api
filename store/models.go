package store

import (
	"time"
	"quadlets-server-api/jobstatus"
)

type User struct {
	Username string
	Password string
	Name string
	Phone string
}

type Bid struct {
	id string
	budget float32
	description string
	createDate time.Time
}

type Job struct {
	id string
	budget float32
	description string
	location Location
	deadline time.Time
	bids []Bid
	acceptedBid *Bid
	status jobstatus.JobStatus
}

type Location struct {
	latitude float64
	logtitude float64
}

type Controller struct {
	Repository Repository
}

type JwtToken struct {
	Token string `json:"token"`
}

type Exception struct {
	Message string `json:"message"`
}
