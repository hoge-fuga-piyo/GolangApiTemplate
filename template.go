package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprint(w, "Hello World")
	})
	http.HandleFunc("/hoge", hogeHandler)
	http.HandleFunc("/post", postHandler)

	http.ListenAndServe(":8080", nil)
}

func hogeHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "hogehoge")
}

func postHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)	// 405
		w.Write([]byte("Only post"))
		return
	}

	w.Write([]byte("OK"))
}
