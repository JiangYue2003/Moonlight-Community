package userfeed

import "time"

// Clock is intentionally minimal so expiry behavior can be tested without
// sleeping while production continues to use wall time.
type Clock interface {
	Now() time.Time
}

type wallClock struct{}

func (wallClock) Now() time.Time { return time.Now() }
