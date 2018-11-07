package main

import "log"

func main() {
	slice1 := make([]string, 5)
	slice2 := []string{"Red", "Green", "Blue"}
	var slice3 []int = nil
	log.Println(slice1)
	log.Println(slice2)
	log.Println(slice3 == nil)
	log.Println(slice3)

	sliceOn()

	for index, value := range slice2 {
		log.Printf("Index: %d, Value: %s\n", index, value)
	}
	for index := 2; index < len(slice2); index++ {
		log.Printf("Index: %d, Value: %s\n", index, slice2[index])
	}
}

func sliceOn() {
	log.Println("切切片：")
	slice := []int{1, 2, 3, 4, 5}
	newSlice := slice[1:3]
	log.Printf("%+v\n", slice)
	log.Printf("%+v：%d\n", newSlice, cap(newSlice))
}

