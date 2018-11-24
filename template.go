package main

import (
	"fmt"
	"net/http"
	"encoding/json"
)

func main() {
	// クロージャを使うパターン
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
	})

	// クロージャを使わないパターン
	http.HandleFunc("/hoge", hogeHandler)

	// POSTリクエストのみ受け付ける
	http.HandleFunc("/post", postHandler)

	// GETリクエストのみ受け付ける
	http.HandleFunc("/get", getHandler)

	// リクエストパラメータを取得する
	http.HandleFunc("/request_params", requestParamsHandler)

	// mapを使ってjson形式のレスポンスを返す
	http.HandleFunc("/json/map", jsonWithMapHandler)

	// 構造体を使ってjson形式のレスポンスを返す
	http.HandleFunc("/json/struct", jsonWithStructHandler)

	http.ListenAndServe(":8080", nil)
}

func hogeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hogehoge")
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)	// 405
		w.Write([]byte("Only post"))
		return
	}

	w.Write([]byte("POST OK"))
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)	// 405
		w.Write([]byte("Only get"))
		return
	}

	w.Write([]byte("GET OK"))
}

func requestParamsHandler(w http.ResponseWriter, r *http.Request) {
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

func jsonWithMapHandler(w http.ResponseWriter, r *http.Request) {
	hash := make(map[string]string)
	hash["hoge"] = "hogehoge"
	hash["fuga"] = "fugafuga"

	json.NewEncoder(w).Encode(hash)
}

type JsonStruct struct {
	Hoge string	`json:"hoge"` 
	Fuga int	`json:"fuga"`
	Piyo bool	`json:"piyo"`
}

func jsonWithStructHandler(w http.ResponseWriter, r *http.Request) {
	st := JsonStruct{}
	st.Hoge = "hogehoge"
	st.Fuga = 100
	st.Piyo = true

	json.NewEncoder(w).Encode(st)
}