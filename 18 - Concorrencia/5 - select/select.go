package main

import (
	"log"
	"time"
)

func main() {
	channel1, channel2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Second / 2)
			channel1 <- "Hello 1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			channel2 <- "Hello 2"
		}
	}()

	for {

		select {
		case msg1 := <-channel1:
			log.Println(msg1)
		case msg2 := <-channel2:
			log.Println(msg2)
		}
	}
}
