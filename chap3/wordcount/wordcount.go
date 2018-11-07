// 单词数量
package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		log.Println("参数错误")
		return
	}

	filename := os.Args[1]
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println(err)
		return
	}

	text := string(contents)
	count := CountWords(text)
	log.Printf("word count:%d\n", count)



}

// CountWords counts the number of words in the specified
// string and returns the count.
func CountWords(text string) (count int) {
	count = len(strings.Fields(text))
	return
}
