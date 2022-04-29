package main

import (
	"fmt"
	"net/http"
)

//this is what server side rendereing is all about
func index(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "<h1>Hello World</h1>")
}

func about(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "<p>This is the about me page</p>")
}

func main(){
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	fmt.Println("Server starting...")
	http.ListenAndServe(":3000", nil)
}