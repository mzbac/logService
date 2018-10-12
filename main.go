package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/mzbac/logService/def"
)

var (
	maxWorker       int
	maxQueue        int
	simpleScheduler SimpleScheduler
)

func init() {
	flag.IntVar(&maxWorker, "w", 30, "The number of workers to start")
	flag.IntVar(&maxQueue, "q", 300, "The size of job queue")
}
func payloadHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	simpleScheduler.Submit(def.Job{})

	w.WriteHeader(http.StatusOK)
}
func main() {
	flag.Parse()
	log.Printf("main start with %d worker, max queue size %d", maxWorker, maxQueue)
	simpleScheduler = SimpleScheduler{}
	dispatcher := Dispatcher{
		MaxWorkers: maxWorker,
		MaxQueue:   maxQueue,
		Scheduler:  &simpleScheduler,
	}
	dispatcher.Run()
	http.HandleFunc("/payload/", payloadHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("starting listening for payload messages")
	} else {
		log.Fatalf("an error occured while starting payload server %s", err.Error())
	}
}
