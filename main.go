package main

import (
	"pooling-project/pooling"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go pooling.RunServer()
	wg.Add(1)
	go pooling.RunWatcher()
	wg.Wait()
}
