package main

import (
	"github.com/mzbac/logService/def"
)

type Scheduler interface {
	Submit(def.Job)
	ConfigureWorkerChan(chan def.Job)
}

type SimpleScheduler struct {
	workerChan chan def.Job
}

func (s *SimpleScheduler) Submit(r def.Job) {
	s.workerChan <- r
}

func (s *SimpleScheduler) ConfigureWorkerChan(c chan def.Job) {
	s.workerChan = c
}
