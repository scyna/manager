package model

import (
	"strconv"
	"time"
)

type Session struct {
	Module      string
	Id          int64
	LastUpdated time.Time
	State       int
	StartTime   time.Time
	EndTime     time.Time
}

func (s Session) String() string {
	return strconv.FormatInt(s.Id, 10)
}
