package main

import "fmt"

func main() {
	x := [5]int{10, 20, 30, 40, 50}
	//
	for i := 0; i < len(x); i++ {
		fmt.Println(x[i])

	}

}
