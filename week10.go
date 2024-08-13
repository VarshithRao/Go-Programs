package main

import "fmt"

func removeDuplicates(arr []int) []int {
	encountered := map[int]bool{}
	result := []int{}

	for _, v := range arr {
		if !encountered[v] {
			encountered[v] = true
			result = append(result, v)
		}
	}

	return result
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 1, 6, 3}
	fmt.Println("Original Array:", arr)
	arr = removeDuplicates(arr)
	fmt.Println("Array after removing duplicates:", arr)
}
