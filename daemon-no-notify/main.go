package main

import (
	"fmt"
	"time"

	"github.com/okzk/sdnotify"
)

func main() {
	loop := true
	idx := -1
	for loop {
		idx += 1
		fmt.Println("==== test-watchdog.watchdog-no-notify running. idx:", idx)
		if idx == 3 {
			//sdnotify.SdNotify("WATCHDOG=1")
			fmt.Println("==== test-daemon NOT notified on watchdog")
		}
		if idx == 10 {
			break
		}
		time.Sleep(1 * time.Second)
	}
		sdnotify.SdNotify("notifying")
}
