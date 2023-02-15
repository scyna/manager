package model

import (
	"strconv"
	"time"
)

type Session struct {
	id          int64
	lastUpdated time.Time
	state       int
}

func (s Session) String() string {
	return strconv.FormatInt(s.id, 10)
}
