package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func say(s string) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		time.Sleep(time.Millisecond * 100)
		fmt.Println(s)
	}
}

func main() {
	wg.Add(1)
	go say("Hey")
	wg.Add(1)
	go say("There")
	// wg.Wait()

	wg.Add(1)
	go say("Hi")
	wg.Wait()
}
