package main

import "log"

func main() {
	ip := []byte{1, 2, 3}
	change(ip)
	log.Println(ip)

	ip = append(ip, ip...) // 解构
	log.Println(ip)
}

func change(num []byte) []byte {
	num[1] = 100
	return num
}
