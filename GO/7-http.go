package main
import (
	"net/http"
	"io"
)
func main() {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8000", nil)
}

func hello(response http.ResponseWriter, request *http.Request){
	io.WriteString(response, "<h2>hello world</h2>")
}

func home_page(response http.ResponseWriter, request *http.Request){
	http.ServeFile(response, request, "index.html")
}
