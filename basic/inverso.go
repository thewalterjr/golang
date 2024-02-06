package main

import "fmt"

func reverse(arr []int) {
	if len(arr) == 0 {
		return
	}

	fmt.Println(arr[len(arr)-1])

	reverse(arr[:len(arr)-1])
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	reverse(numbers)
}
