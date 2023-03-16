package main

import (
	"fmt"
	"strings"
)

func parseTest(sentences []string, chars []rune) [][]int {
	var strSlice []string
	var res [][]int

	for _, sentence := range sentences {
		splitedSent := strings.Split(sentence, " ")
		strSlice = append(strSlice, splitedSent[len(splitedSent)-1])
	}
	for _, word := range strSlice {
		var tempRes []int
		for _, char := range chars {
			tempRes = append(tempRes, strings.Index(strings.ToTitle(word), strings.ToTitle(string(char))))
		}
		res = append(res, tempRes)
	}
	return res
}
func main() {
	sentences := []string{"Hello world", "Hello Skillbox", "Привет Мир", "Привет Skillbox"}

	chars := []rune{'h', 'E', 'l', 'П', 'М'}

	res := parseTest(sentences, chars)
	fmt.Println(res)

}
