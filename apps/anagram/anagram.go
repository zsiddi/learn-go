package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("testing anagram")

	var s1 = "Debit card "
	var s2 = "Bad credit "

	fmt.Printf("Are  %s and %s Anagram = %t ", s1, s2, areAnagram(s1, s2))
}

func areAnagram(str1 string, str2 string) bool {

	//Step 1: Replace all spaces and bring string to same case

	str1 = strings.ReplaceAll(str1, " ", "")
	str2 = strings.ReplaceAll(str2, " ", "")

	s1 := strings.ToLower(str1)
	s2 := strings.ToLower(str2)

	//Step 2: check if the lengths are same
	if len(s1) != len(s2) {
		return false
	}

	/*
		store both string in hashmap using a number, increment is already seen in the map.
		If anagram, all the character in string should have the value 2
	*/

	hashmap := make(map[string]int)

	for _, v := range s1 {
		j := hashmap[string(v)]

		if j == 0 {
			hashmap[string(v)] = 1
		} else {
			hashmap[string(v)] = 1 + 1
		}
	}

	for _, v := range s2 {
		j := hashmap[string(v)]

		if j == 0 {
			hashmap[string(v)] = 1
		} else {
			hashmap[string(v)] = 1 + 1
		}
	}

	var isAnagram bool = true
	for _, v := range hashmap {
		if v%2 != 0 {
			isAnagram = false
			break
		}
	}
	return isAnagram
}
