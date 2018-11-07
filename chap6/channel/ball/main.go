package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}
func main() {
	court := make(chan int)
	wg.Add(2)

	// 启动两个选手
	go player("Bill", court)
	go player("Gate", court)

	court <- 1 // 发球

	wg.Wait()
}

func player(name string, court chan int) {
	defer wg.Done()
	for {
		// 等待球被击打过来
		ball, ok := <-court
		if !ok {
			// 如果通道被关闭，我们就赢了
			fmt.Printf("Player %s Won.\n", name)
			return
		}

		// 选随机数，然后用这个数来判断我们是否丢球
		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("Player %s Missed.\n", name)
			// 关闭通道，表示我们输了
			close(court)
			return
		}

		// 显示击球数，并将击球数加 1
		fmt.Printf("Player %s Hit %d\n", name, ball)
		ball++ // 将球打向对手

		court <- ball
	}
}
