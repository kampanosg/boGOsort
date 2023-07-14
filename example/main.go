package main

import (
	"log"
	"time"

	bogosort "github.com/kampanosg/boGOsort"
)

func main() {
	start := time.Now()
	s := []int {2,3,4,1}

	// Please don't use this :')
	bogosort.Sort(s)

	log.Printf("done in %v", time.Since(start))
	log.Printf("result: %v", s)
}
