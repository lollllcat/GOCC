package main

import (
	"sync"
)

func main() {
	sle := sync.Mutex{}
	sle.Lock()
	defer sle.Unlock()

	go func() {
		for {
			select {
			default:
				sle.Lock()
				sle.Unlock()
			}
		}
	}()
}
