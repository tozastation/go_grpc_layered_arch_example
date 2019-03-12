package timer

import (
	"fmt"
	"github.com/tozastation/go-grpc-ddd-example/interfaces/timer/itimer"
	"time"
)

// Timer is .
type Timer struct{}

// NewTimer is .
func NewTimer() itimer.ITimer {
	return &Timer{}
}

// Start is .
func (timer *Timer) Start() time.Time {
	return time.Now()
}

// End is .
func (timer *Timer) End() time.Time {
	return time.Now()
}

// Result is .
func (timer *Timer) Result(ctx string, start, end time.Time) {
	fmt.Println("------------------------------------------------------------------------------------------------")
	fmt.Println("Execute:", start)
	fmt.Println("Query:", ctx)
	fmt.Println("Result:", end.Sub(start).Seconds())
	fmt.Println("------------------------------------------------------------------------------------------------")
}
