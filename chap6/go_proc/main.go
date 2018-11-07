package main

import "fmt"

func main() {
	go task(1)
	go task(2)
	go task(3)
	go task(4)

	select {}

}

func task(i int) {

	for {
		fmt.Println(i)
	}
}
