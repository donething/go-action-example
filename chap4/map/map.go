package main

import (
	"log"
)

func main() {
	dict := make(map[string]int)
	colors := map[string]string{"Red": "#da1337", "Orange": "#e95a22"}
	log.Println(dict, colors)

	for key, value := range colors {
		log.Printf("Key: %s, Value: %s\n", key, value)
	}
	var m map[string]string
	m = map[string]string{}
	m = make(map[string]string)
	m["key"] = "value"
	log.Printf("%v\n", m)

	useMap()
}

func useMap() {
	colors := map[string]string{"Red": "#da1337", "Orange": "#e95a22"}
	value, exists := colors["Bule"]
	if exists {
		log.Println(value)
	} else {
		log.Println("键不存在")
	}

	value = colors["Red"]
	if value != "" {
		log.Println(value)
	} else {
		log.Println("键不存在")
	}
}
