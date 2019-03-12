package itimer

import "time"

// ITimer is .
type ITimer interface {
	Start() time.Time
	End() time.Time
	Result(ctx string, start, end time.Time)
}
