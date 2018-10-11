package main

import (
	"github.com/mzbac/logService/def"
)

type Dispatcher struct {
	Scheduler  Scheduler
	maxWorkers int
}

func (d *Dispatcher) Run() {
	in := make(chan def.Job)
	d.Scheduler.ConfigureWorkerChan(in)
	for i := 0; i < d.maxWorkers; i++ {
		createWorker(in, d.Scheduler)
	}
}

func createWorker(in chan def.Job, scheduler Scheduler) {
	go func() {
		for {
			r := <-in
			_, err := Worker(r)
			if err != nil {
				scheduler.Submit(r)
				continue
			}
		}
	}()
}