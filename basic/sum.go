package main

import "fmt"

func soma(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	return arr[0] + soma(arr[1:])
}

func main() {
	arr := []int{2, 4, 6}
	fmt.Println("Res:", soma(arr))
}
