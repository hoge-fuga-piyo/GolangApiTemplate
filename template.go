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
	http.HandleFunc("/request_params", requestParamsHandler)

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

func requestParamsHandler(w http.ResponseWriter, r *http.Request){
	// クエリパラメータの取得
	fmt.Fprintf(w, "Query:%s\n", r.URL.RawQuery)

	// Bodyデータを扱う場合は事前にパースが必要
	r.ParseForm()

	// Formデータの取得
	form := r.PostForm
	fmt.Fprintf(w, "Form:\n%v\n", form)

	// クエリパラメータ、Formデータの両方
	params := r.Form
	fmt.Fprintf(w, "Form2:\n%v\n", params)
}
