package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func main() {

	var start = time.Now()

	var waitGroup sync.WaitGroup

	waitGroup.Add(2)

	go func() {
		write("Hello")
		waitGroup.Done()
	}()

	go func() {
		write("Programando em Go")
		waitGroup.Done()
	}()

	waitGroup.Wait()

	log.Println("Tempo: ", time.Now().Sub(start))

}

func write(text string) {
	for i := 0; i < 5; i++ {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
