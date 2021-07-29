package main

import "fmt"

func main() {
	var arr = []int{5, 6, 13, 6, 9}

	fmt.Printf("the largest element is %d \n", findLargetElement(arr))
}

func findLargetElement(ar []int) int {
	var l int

	for i := 0; i < len(ar); i++ {
		if l <= ar[i] {
			l = ar[i]
		}
	}
	return l
}
