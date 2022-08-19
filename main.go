package main

import "fmt"

func main() {
	arr := []int{3, 2, 3}
	result := majorityElement(arr)

	fmt.Println(result)
}

func majorityElement(nums []int) int {
	nums = sort(nums)
	mid := len(nums) / 2

	return nums[mid]
}

func sort(arr []int) []int {

	for i := 0; i < len(arr); i++ {

		for j := 0; j < len(arr); j++ {

			if i != j && arr[i] < arr[j] {
				arr = swap(arr, i, j)
			}

		}
	}
	return arr
}

func swap(arr []int, i, j int) []int {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp

	return arr
}
