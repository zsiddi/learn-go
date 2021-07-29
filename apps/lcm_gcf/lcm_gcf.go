package main

import "fmt"

func main() {

	i, j := 12, 18

	lcm := lcm(i, j)
	gcf := gcf(i, j)
	fmt.Printf("The LCM for %v and %v is - %v \n", i, j, lcm)
	fmt.Printf("The GCF for %v and %v is - %v \n", i, j, gcf)
}

func lcm(i, j int) int {

	var lcm int

	if i > j {
		lcm = i
	} else {
		lcm = j
	}

	var in int = 2

	for {
		lcm = lcm * in

		if lcm%i == 0 && lcm%j == 0 {
			break
		} else {
			in++
		}
	}
	return lcm
}

func gcf(i, j int) int {
	var gcf int
	for in := 2; in <= i && in <= j; in++ {
		if i%in == 0 && j%in == 0 {
			gcf = in
		}
	}
	return gcf
}
