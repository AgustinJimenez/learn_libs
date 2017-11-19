package main
import (
	"fmt"
)
func main() {
	var v int
	fmt.Println( "Select value" )
	fmt.Scanf("%d", v);
	test(v)
}

func test(value int){

	if value >= 1 {
		fmt.Println( "1" )
	}

	if value >= 2 {
		fmt.Println( "2" )
	}

	if value >= 3 {
		fmt.Println( "3" )
	}

	if value >= 4 {
		fmt.Println( "3" )
	}

	

}