package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println("working...")
		time.Sleep(10 * time.Second)
	}
}
