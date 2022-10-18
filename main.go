package main

import (
	"net/http"
	"os"
	"sync"
)

func main() {
	kvService := &KVService{}
	go kvService.InitKVService()
	go InitTestService()
	err := http.ListenAndServe(":10000", nil)
	if err != nil {
		os.Exit(2)
	}
	wg := sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()
}
