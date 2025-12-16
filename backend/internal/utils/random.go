package utils

import (
	"math/rand"
	"time"
)

// RandomSleep sleeps for a random duration between min and max milliseconds
func RandomSleep(min, max int) {
	duration := time.Duration(rand.Intn(max-min)+min) * time.Millisecond
	time.Sleep(duration)
}
