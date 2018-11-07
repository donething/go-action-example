package main

import "log"

func main() {
	// 数组是值类型
	array := [5]int{1, 2, 3, 4, 5}
	log.Println(array)
	modArray(array)
	log.Println(array)

	// 切片是引用类型
	slice := []int{1, 2, 3, 4, 5}
	log.Println(slice)
	modSlice(slice)
	log.Println(slice)

	// 数组指针的赋值
	log.Println("数组指针的赋值")
	arrayPointerCopy()
}

func modArray(array [5]int) {
	array[2] = 1000
}

func modSlice(slice []int) {
	slice[2] = 1000
}

func arrayPointerCopy() {
	var array1 [3]*string
	var array2 = [3]*string{new(string), new(string), new(string)}
	*array2[0] = "Red"
	*array2[1] = "Green"
	*array2[2] = "Blue"

	log.Printf("array1: %+v", array1)
	log.Printf("array1： %+v", array2)
	array1 = array2 // 指针数组赋值，指向了相同的值
	log.Printf("array1: %+v", array1)
	log.Printf("array1： %+v", array2)

	log.Println(1e1)
}
