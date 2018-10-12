package main

import (
	"log"

	"github.com/mzbac/logService/def"
)

func Worker(r def.Job, idx int) (int, error) {
	// TODO
	log.Printf("worker %d takes new job", idx)
	return 0, nil
}
