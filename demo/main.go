package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("You can view the application at https://localhost:8080/ CTRL+C to stop the server.")
	_ = http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))
}
