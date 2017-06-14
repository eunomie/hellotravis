package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!!!"))

	fmt.Fprintf(w, "\nEnvironment:\n")
	for _, env := range os.Environ() {
		fmt.Fprintf(w, "%s\n", env)
	}
	fmt.Fprintf(w, "\n")

	hostname, err := os.Hostname()
	if err != nil {
		fmt.Fprintf(w, "Error: %s\n", err)
		return
	}
	fmt.Fprintf(w, "Generated on %s\n", hostname)
}

func main() {
	fmt.Println("hello world")
	http.HandleFunc("/", HelloWorld)
	log.Fatal(http.ListenAndServe(":80", nil))
}
