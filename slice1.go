package main

import "fmt"

func advanceOfSlice1() {

	array := []int{10, 5, 12, 15, 20, 25, 18}

	s := array[1:5]

	s[1] = 23
	fmt.Println(s[1])
	fmt.Println(array)
}
