package main

import "fmt"

func main() {
	var x [10]int
	fmt.Println(x)
	fmt.Println(len(x))

	x[0] = 10
	x[1] = 6
	x[2] = 12

	fmt.Println(x)

	z := [5]int{5, 10, 4, 5, 6}

	fmt.Println(z)

}
