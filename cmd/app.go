package main

import (
	"math/rand"
	"time"
	"github.com/primat/goker"
)

// Application entry point
func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	goker.Profile()
}
