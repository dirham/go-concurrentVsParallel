package main

import (
	"conendpar/concur"
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(1) // ini concurrent
	// runtime.GOMAXPROCS(3) // ini parallel
	var tunggu sync.WaitGroup
	tunggu.Add(3)

	go concur.Ikatsepatu(&tunggu)
	go concur.Lari(&tunggu)
	go concur.Dengarmusik(&tunggu)
	fmt.Println("--------------------------------------------------")
	tunggu.Wait()
}
