/*
 * @Author: your name
 * @Date: 2021-08-02 15:55:05
 * @LastEditTime: 2021-08-02 16:09:51
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \project\Handle\handle.go
 */
package main

import (
	"fmt"
	"net/http"
)

type hellohandler struct{}

func (m *hellohandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello!")
}
func Home(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Home Te!")
}
func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "About!")
}
func main() {
	mh := hellohandler{}

	http.Handle("/hello", &mh)

	http.HandleFunc("/welcome", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(rw, "Welcome!")
	})
	http.HandleFunc("/home", Home)

	http.Handle("/about", http.HandlerFunc(about))

	http.ListenAndServe(":8080", nil)
}
