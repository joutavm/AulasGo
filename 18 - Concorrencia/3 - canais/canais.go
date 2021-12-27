package main

import (
	"log"
	"time"
)

func main() {

	var start = time.Now()

	canal := make(chan string)

	go write("Olá mundo", canal)

	//for {
	//	mensagem, aberto := <-canal
	//	if !aberto {
	//		break
	//	}
	//	log.Println(mensagem)
	//}

	for message := range canal {
		log.Println(message)
	}

	log.Println("Tempo: ", time.Now().Sub(start))

}

func write(text string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- text
		time.Sleep(time.Second)
	}

	close(canal)
}
