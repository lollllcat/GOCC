package main

import (
	"sync"
)

func main() {
	sle := sync.Mutex{}
	go func() {
		go func() {
			for {
				select {
				default:
					sle.Lock()
					sle.Unlock()
				}
			}
		}()
	}()
}
