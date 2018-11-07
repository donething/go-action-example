package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	ch := make(chan int, 2)
	wg.Add(1)

	ch <- 1
	fmt.Println("开始：")
	close(ch)

	for n := range ch {
		fmt.Printf("当前值：%d", n)
	}
	
}
