package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi %s!", r.URL.Path[1:])
}

func main() {
	//アクセスのルーティングを設定
	http.HandleFunc("/", handler)
	//監視するポートを設定
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}