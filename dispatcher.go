package main

import (
	"github.com/mzbac/logService/def"
)

type Dispatcher struct {
	Scheduler  Scheduler
	MaxWorkers int
	MaxQueue   int
}

func (d *Dispatcher) Run() {
	in := make(chan def.Job, d.MaxQueue)
	d.Scheduler.ConfigureWorkerChan(in)
	for i := 0; i < d.MaxWorkers; i++ {
		createWorker(in, d.Scheduler, i)
	}
}

func createWorker(in chan def.Job, scheduler Scheduler, index int) {
	go func() {
		for {
			r := <-in
			_, err := Worker(r, index)
			if err != nil {
				scheduler.Submit(r)
				continue
			}
		}
	}()
}
