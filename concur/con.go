package concur

import (
	"fmt"
	"sync"
	"time"
)

func Lari(wg *sync.WaitGroup) {
	time.Sleep(time.Second * 2)
	fmt.Println("lari Sore...", time.Now().Format("15:04:05"))
	wg.Done()
}

func Ikatsepatu(wg *sync.WaitGroup) {
	time.Sleep(time.Second * 2)

	fmt.Println("Ikat sepatu...", time.Now().Format("15:04:05"))
	wg.Done()
}

func Dengarmusik(wg *sync.WaitGroup) {
	time.Sleep(time.Second * 2)
	fmt.Println("Dengar musrik...", time.Now().Format("15:04:05"))
	wg.Done()
}
