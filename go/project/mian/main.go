/*
 * @Author: your name
 * @Date: 2021-06-08 22:55:33
 * @LastEditTime: 2022-01-04 21:47:06
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \Go\src\project\mian\main.go
 */
package main

import (
	"fmt"
	"log"
	"net/http"
)

func dy(in interface{}) {
	if v, ok := in.(string); ok {
		fmt.Println(v)
	}
}

type hellohandler struct{}

func (m *hellohandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}

type welcomehandler struct{}

func (m *welcomehandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome!"))
}
func about(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("About!"))
	fmt.Fprintln(w, "About!")
}

func main() {
	var tslice []int
	tslice = make([]int, 10)
	var s string = "HUUdsf"
	dy(s)
	fmt.Println(tslice)
	fmt.Println(log.Ltime)
	mh := hellohandler{}
	wl := welcomehandler{}
	service := http.Server{
		Addr:    "localhost:8080",
		Handler: nil,
	}
	http.Handle("/hello", &mh)
	http.Handle("/welcome", &wl)
	http.Handle("/about", http.HandlerFunc(about))
	service.ListenAndServe()
	http.HandleFunc("/hello", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("Hello!"))
		fmt.Fprintln(rw, "LLLLL")
	})
	// http.Handler
	// http.HandlerFunc("")
	// http.DefaultServeMux
	// http.ListenAndServe(":808", nil)

}
