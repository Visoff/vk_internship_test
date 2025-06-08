package workerpool

import (
	"fmt"
	"sync"
	"time"
)

type WorkerPool struct {
	ch      chan string
	quit    chan struct{}
	workers []int
	wg      *sync.WaitGroup
}

func NewWorkerPool(size int) *WorkerPool {
	pool := &WorkerPool{
		ch:      make(chan string),
		quit:    make(chan struct{}),
		workers: make([]int, size),
		wg:      &sync.WaitGroup{},
	}

	for i := range size {
		pool.wg.Add(1)
		pool.workers[i] = i
		go pool.worker(i)
	}

	return pool
}

func (p *WorkerPool) process(msg string, id int) {
	//time.Sleep(time.Second)
	fmt.Println(msg, id)
}

func (p *WorkerPool) worker(id int) {
	defer p.wg.Done()
	for {
		select {
		case <-p.quit:
			return
		case msg := <-p.ch:
			p.process(msg, id)
		}
	}
}

func (p *WorkerPool) Stop() {
	close(p.quit)
	p.wg.Wait()
}

func (p *WorkerPool) AddWorkers(num int) {
	for range num {
		p.wg.Add(1)
		go p.worker(len(p.workers))
		p.workers = append(p.workers, len(p.workers))
	}
}

func (p *WorkerPool) RemoveWorkers(num int) {
	p.quit <- struct{}{}
}

func (p *WorkerPool) Process(msg string) {
	p.ch <- msg
}
