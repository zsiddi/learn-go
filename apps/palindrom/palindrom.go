package main

import "fmt"

func main() {
	var str = "113"

	fmt.Printf("is %s a Palindrom - %t", str, isPalindrom(str))
}

func isPalindrom(s string) bool {

	for i, j := 0, (len(s) - 1); i <= j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}
