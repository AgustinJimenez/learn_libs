package main
import (
	"fmt"
)
func main() {
	var v int

	fmt.Println( "Select value" )
	fmt.Scanf("%d", v)
	if test(v){
		fmt.Println("return true\n\n")
	}else{
		fmt.Println("return false\n\n")
	}


}

func test(value int) bool{

	defer func(){
		fmt.Println("DEFER HERE before return always!!!")
	}()

	fmt.Println( " begin function \n" )

	switch value{
	case 1:
		fmt.Println( "is 1" )
		return true
	case 2:
		fmt.Println( "is 2" )
		return true
	default:
		fmt.Println( "none" )
		return false
	}

	fmt.Println( " end function \n\n" )
	return false
}