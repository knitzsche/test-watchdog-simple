package main

import (
	"fmt"
	"time"
)

func main() {
	for { // loop forever
		time.Sleep(5 * time.Second)
		fmt.Println("==== test-daemon running\n")
	}
}
