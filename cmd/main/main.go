package main

import (
	"github.com/Visoff/vk_internship_test/pkg/worker_pool"
)

func main() {
	pool := workerpool.NewWorkerPool(5)
	pool.AddWorkers(5)
	for range 10 {
		pool.Process("message")
	}
	pool.Stop()
}
