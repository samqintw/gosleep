package main

import (
	"flag"
	"log"
	"time"
)

var (
	times = flag.Int("t", 10, "in Minutes")
)

func main() {
	flag.Parse()
	log.Printf("falling into sleep %d minutes", *times)
	time.Sleep(time.Minute * time.Duration(*times))
}