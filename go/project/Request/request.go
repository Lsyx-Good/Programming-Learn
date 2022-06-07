/*
 * @Author: your name
 * @Date: 2021-08-02 16:36:25
 * @LastEditTime: 2022-01-06 17:47:27
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \project\Request\request.go
 */
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/url", http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(rw, r.URL.Fragment)
	}))

	http.Handle("/header", http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(rw, r.Header)
		fmt.Fprintln(rw, r.Header["Accept-Encoding"])
		fmt.Fprintln(rw, r.Header.Get("Accept-Encoding"))
	}))
	http.Handle("/body", http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		lenght := r.ContentLength
		body := make([]byte, lenght)
		r.Body.Read(body)
		fmt.Fprintln(rw, string(body))
	}))
	http.Handle("/query", http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(rw, r.URL.Query())
	}))
	http.ListenAndServe(":8080", nil)
}
