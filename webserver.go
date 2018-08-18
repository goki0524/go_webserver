package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	//オプションを解析
	r.ParseForm()
	//データはサーバのプリント情報に出力
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	//wに入るものはクライアントに出力
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
