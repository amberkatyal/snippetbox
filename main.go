package main

import (
	"fmt"
	"log"
	"net/http"
)

// Define a home handler function which writes a byte slice containing
// "Hello from Snippetbox!" as the response body.
//
// *What is a ResponseWriter?
// It is an interface type that allows us to construct an HTTP response which is used by the http handler.
// It has methods for adding headers, setting cookies, and writing the response body. It is
// passed to the handler by the http server.
//
// *What is a Request? It is a struct type that holds all the information
// about the HTTP request(http method, url, headers, body etc) that was sent by the client.
// Why it is a pointer and not a class?
// Because there are no classes in Go, and using a pointer allows us to avoid copying the entire struct.
func home(w http.ResponseWriter, r *http.Request) {
	// In Swift we use type Data, here we use byte slice.
	// It is similar to Swift's Data type which represents a contiguous block of memory.
	// What is byte slice?
	// A Slice is a data structure in Go that represents a dynamic array, which is represented by Array in Swift.
	// Whereas an Array in Go is a fixed-size sequence of elements of a specific type.
	// A Slice is declared using the syntax []sliceType and array as [arraySize]arrayType.
	// A byte slice is a slice whose elements are of type byte.
	// A byte is an alias for uint8, which means it is an unsigned 8-bit integer.
	// A byte slice is commonly used to represent binary data or text data in Go and single byte can represent a character.

	// Here we are converting the string "Hello from Snippetbox!" to a byte slice using []byte() conversion.
	// Then we write it to the ResponseWriter using the Write method.
	// The Write method writes the data to the connection as part of an HTTP reply.
	// It returns the number of bytes written and any error encountered.
	bytes, error := w.Write([]byte("Hello from Amber!"))
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println("Bytes written:", bytes)
	}
}

func main() {
	// What is a router in web development?
	// A router is a component that maps incoming HTTP requests to specific handler functions based on the request's URL path and HTTP method.
	// In Go's net/http package, the default router is used when we call http.HandleFunc.
	// But we can use servemux too. What is servemux?
	// A ServeMux (short for "HTTP request multiplexer") is an HTTP request router provided by the net/http package in Go.

	mux := http.NewServeMux()
	// Register the home handler function for the "/" URL path.
	mux.HandleFunc("/", home)

	log.Println("starting server on: 4000")

	// Start the HTTP server on port 4000 and use the mux as the handler.
	// The function takes two parameters: the TCP network address to listen on (in this case, ":4000" which means all interfaces on port 4000)
	// and the handler to use (in this case, mux).
	// This is important to note that ListenAndServe is a blocking function, meaning it will run indefinitely until the program is terminated.
	// This will always return a non-nil error which means the server has stopped or failed at any stage during starting.
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
