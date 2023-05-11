package main

import (
	"flag"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {
	durationSecond := flag.Int("duration", 0, "duration of program need to run, unit is second")
	offset := flag.Int("offset", 0, "offset of duration, unit is second")
	flag.Parse()
	log.SetPrefix("[perfCPU]")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile | log.Lmsgprefix)
	if *durationSecond <= 0 {
		log.Println("duration of program must greater than 0")
		os.Exit(1)
	}
	if *offset < 0 {
		log.Println("offset of duration cat not less than 0")
		os.Exit(1)
	}

	durValue := *durationSecond - *offset + rand.New(rand.NewSource(time.Now().UnixNano())).Intn(*offset*2)
	log.Println("duration of program run is:", durValue)

	exitChan := make(chan struct{}, 1)
	ackChan := make(chan struct{}, 1)
	go calculate(exitChan, ackChan)
	time.Sleep(time.Duration(durValue) * time.Second)
	exitChan <- struct{}{}
	close(exitChan)
	_, ok := <-ackChan
	if ok {
		log.Println("receive ack msg from calculate thread")
	}
	log.Println("main thread exit with code 0")
	os.Exit(0)
}

func calculate(exitChan <-chan struct{}, ackChan chan<- struct{}) {
	var (
		v  float64
		i  float64
		ok bool
	)
	log.Println("start calculate thread")
	for {
		select {
		case _, ok = <-exitChan:
			if ok {
				log.Println("receive exit msg from main thread, send ack msg to main thread")
				ackChan <- struct{}{}
				close(ackChan)
				log.Println("exit calculate thread")
				return
			}
			log.Println("exit channel is closed, no msg")
		default:
			// log.Println("default case has been selected")
		}

		for i = 0; i < 1000; i += 1 {
			v = 3.14 * 2 * i
			v = v / 6.28
		}
	}
}
