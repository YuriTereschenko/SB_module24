package main

import "fmt"

func spreadArr(input []int) ([]int, []int) {
	var odd, even []int
	for _, value := range input {
		if value%2 == 0 {
			even = append(even, value)
		} else {
			odd = append(odd, value)
		}
	}
	return odd, even
}
func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	odd, even := spreadArr(arr)
	fmt.Println("odd is:", odd)
	fmt.Println("even is:", even)
}
