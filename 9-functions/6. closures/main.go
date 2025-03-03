package main

import (
	"fmt"
	"time"
)

func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func rateLimiter(maxCalls int, duration time.Duration) func() bool {
	count := 0
	lastReset := time.Now()

	return func() bool {
		if time.Since(lastReset) > duration {
			count = 0
			lastReset = time.Now()
		}
		if count < maxCalls {
			count++
			return true
		}
		return false
	}
}

func main() {
	increment := counter()
	fmt.Println(increment())
	fmt.Println(increment())

	// Usage: Allow 3 requests/second
	checkLimit := rateLimiter(3, time.Second)
	fmt.Println(checkLimit())

}
