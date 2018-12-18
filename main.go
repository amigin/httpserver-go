package main

import "fmt"
import "net/http"

func main() {
	http.Handle("/", http.FileServer(http.Dir("./wwwroot")))
	fmt.Println("Listening *:5000")
	http.ListenAndServe(":5000", nil)
}
