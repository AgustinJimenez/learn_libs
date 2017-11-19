package main

import "fmt"

func main(){
	a := 10
	LOOP_LABEL: for a < 20{
		if a == 15 {
			a++
			goto LOOP_LABEL
		}
		fmt.Printf("Value of a: %d\n", a)
		a++
	}

}