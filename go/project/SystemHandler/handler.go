/*
 * @Author: your name
 * @Date: 2021-08-05 18:50:46
 * @LastEditTime: 2021-08-05 22:52:34
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \project\SystemHandler\handler.go
 */
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/url", http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("Hello Redir\n"))
		fmt.Fprintln(rw, http.StatusMovedPermanently)
	}))
	http.Handle("/handle", http.RedirectHandler("/url", http.StatusMovedPermanently))

	http.ListenAndServe(":8080", nil)
}
