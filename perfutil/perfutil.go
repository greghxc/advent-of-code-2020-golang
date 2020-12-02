package perfutil

import (
	"fmt"
	"time"
)

// Measure ...
func Measure(times int, f func() string) {
	start := time.Now()
	ans := ""
	for i := 0; i < times; i++ {
		ans = f()
	}
	duration := time.Since(start)
	fmt.Printf("%v (Completed %v times in %v)\n", ans, times, duration)
}
