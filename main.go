package main

import (
	"log"
	"net/http"

	"github.com/mzbac/logService/def"
)

var (
	MaxWorker       = 20 //os.Getenv("MAX_WORKERS")
	simpleScheduler SimpleScheduler
)

func payloadHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	simpleScheduler.Submit(def.Job{})

	w.WriteHeader(http.StatusOK)
}
func main() {
	log.Println("main start")
	simpleScheduler = SimpleScheduler{}
	dispatcher := Dispatcher{
		maxWorkers: MaxWorker,
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
