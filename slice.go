package main

import "fmt"

// creating a slice from existing array
func basicOfSlice() {
	arr := [5]int{50, 30, 40, 20, 35}

	slice := arr[1:4]
	fmt.Println("slice :", slice)
	fmt.Println("len :", len(slice))
	fmt.Println("cap :", cap(slice))
}

// modifying slice changes array
func basicOfSlice1() {

	arr := [5]int{50, 30, 40, 20, 35}

	slice := arr[1:4]

	slice[0] = 200

	fmt.Println("arr:", arr)

}

func basicOfSlice2() {

	s := []int{20, 15, 18}
	fmt.Println("s :", s)
	fmt.Println("len :", len(s))
	fmt.Println("cap :", cap(s))

	fmt.Println("After Modifying ")
	s1 := append(s, 22)
	fmt.Println("s1 :", s1)
	fmt.Println("len :", len(s1))
	fmt.Println("cap :", cap(s1))
}

func basicOfSlice3() {

	s := make([]int, 3)

	s[0] = 10
	s[1] = 20
	s[2] = 30
	s = append(s, 40, 50, 60, 70, 80, 90, 100)
	fmt.Println("s: ", s)
	fmt.Println("Len :", len(s))
	fmt.Println("Cap :", cap(s))
}

// creating a slice from another slice
func basicOfSlice4() {

	s := []int{10, 14, 18}

	s1 := s[1:]

	fmt.Println("s1 :", s1)
}

// Copying Slices
func basicOfSlice5() {
	s1 := []int{10, 20, 30}

	s2 := make([]int, len(s1))

	copy(s2, s1)

	fmt.Println(s2)

}

//Removing an element
func basicOfSlice6() {

	s := []int{10, 20, 30, 40, 50}

	// if i want to remvoe the element of index 2

	i := 2

	s = append(s[:i], s[i+1:]...)

	fmt.Println(s)
}
