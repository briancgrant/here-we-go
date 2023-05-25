package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"

	jn "github.com/jnguyen0220/here-we-go-function/intro"
)

func main() {
	// 1. Hello World
	fmt.Println("Hello World")
	// 2. Custom Function - Use Command: 'go run .' to include package files
	test("There")
	// 3. Third Party Library
	fmt.Println(rand.Intn(100))
	// 4. GitHub Package (1st Party Library)
	jn.Hello()
	// 5. Http Server
	port := 9090
	http.HandleFunc("/", handler)
	fmt.Printf("Http server running on :%v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func handler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Hi\n")
}
