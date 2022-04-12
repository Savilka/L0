package main

import (
	"L0/internal"
)

func main() {
	var a internal.App
	a.Run("test-cluster", "service", "nats://0.0.0.0:4222")
}
