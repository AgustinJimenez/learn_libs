package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println("hello world \n=======================================================================================\n\n\n")

	var one, two int

	var a_string string

	var three = 3

	one = 1

	two = 2

	a_string = "int values to string cocat ( one=" + strconv.Itoa(one) + ", two=" + strconv.Itoa(two) + ", three=" + strconv.Itoa(three) + " ) \n"

	fmt.Println(a_string)

	a_boolean := true //undefined variable

	fmt.Printf("( one=%d, two=%d, three=%d, a_boolean=%d)\n", one, two, three, a_boolean)

	num_string := "is " + "1500"

	num_int, _ := strconv.Atoi(num_string) // atoir return two values, use _ operator to delete second (error) value

	a_string = "\nnum_string=" + num_string + ", num_int=" + strconv.Itoa(num_int)

	fmt.Println(a_string)

	//only loop in go

	for i := 0; i < 10; i++ {

		if i == 4 {
			fmt.Println("five")
			continue
		}
		fmt.Println(i + 1)

	}

	//arrays

	var any_int_array [10]int
	var any_string_array [10]string
	other_int_array := [5]int{1, 2, 3}
	fmt.Println(any_int_array, any_string_array, other_int_array)

	for i := 0; i < len(other_int_array); i++ {

		fmt.Println(i + 1)

	}

	var d2_array [2][3]int

	fmt.Println(d2_array)

	//slices
	var any_slice []int
	fmt.Println(any_slice)
	other_slice := []int{1, 2, 3, 4, 5}
	fmt.Println(other_slice)

	//make and append

	any_slice2 := make([]string, 0, 5)
	any_slice2 = append(any_slice2, "here")
	any_slice2 = append(any_slice2, "here 1")
	any_slice2 = append(any_slice2, "here 2")
	fmt.Println("actual length for elements, max capacity, slice-->", len(any_slice2), cap(any_slice2), any_slice2)

	//copy

	any_slice3 := []int{1, 2, 4, 5}

	any_copy := make([]int, 1, 4)
	copy(any_copy, any_slice3)
	fmt.Println("copy, slice=====>", any_copy, any_slice3)

	fmt.Println("\n\n\n=======================================================================================")

}
