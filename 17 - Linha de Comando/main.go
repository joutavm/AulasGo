package main

import (
	"fmt"
	"linha-de-comando/app"
	"log"
	"os"
)

func main() {
	fmt.Println("Start")

	application := app.Build()
	if err := application.Run(os.Args); err != nil {
		log.Fatalln(err)
	}
}
