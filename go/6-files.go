package main
import (
	"fmt"
	"io/ioutil"
	"bufio"
	"os"
)

func read_file(file_path string, method int ){
	if method == 1{
		read_random_file(file_path)
	}else if method == 2{
		read_random_file_line_by_line(file_path)
	}else{
		fmt.Println( "Invalid method" )
	}
}

func read_random_file(file_path string){
	file_data, err := ioutil.ReadFile(file_path)
	if err!= nil{
		fmt.Println("Hubo un error")
	}else{
		fmt.Println( string(file_data) )
	}
}

func read_random_file_line_by_line(file_path string){
	
	file_data, err := os.Open(file_path)

	defer func(){
		file_data.Close()
		fmt.Println("DEFER ALWAYS CLOSE FILE")
	}()


	if err!= nil{
		fmt.Println("Hubo un error")
	}else{
		scanner := bufio.NewScanner(file_data)
		i := 0
		for scanner.Scan(){
			i++
			line := scanner.Text()
			fmt.Println(i, " - ", line)
		}
	}

	
}

func main() {
	option := 0
	for{
		fmt.Println("\nInsert option (1,2): ")
		fmt.Println("\n======================================>", option)
		_, err := fmt.Scanf("%d", &option)
		if err != nil {
			panic(err)
			recover()
		}else{
			fmt.Println("\n======================================")
			read_file("./hello.txt", option )
			fmt.Println("\n======================================")
		}
		
	}
	
}
