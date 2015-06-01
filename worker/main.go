package main

import (
	"fmt"
	"time"

	"github.com/freeformz/go-heroku-example/message"
)

func main() {
	for {
		fmt.Println(message.Hello)
		time.Sleep(10 * time.Second)
	}
}
