package main

import (
	"fmt"
	"log"
)

type duration int

func (d *duration) pretty() string {
	return fmt.Sprintf("Duration: %d", *d)
}
func main() {
	duration(42).pretty() // 运行错误：cannot take the address of duration(42)
	a := duration(42)
	log.Println(a.pretty())
}
