/*
 * @Author: your name
 * @Date: 2021-06-08 22:51:33
 * @LastEditTime: 2021-08-02 15:57:13
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \project\server\server.go
 */
package main

import (
	"fmt"
	"net/http"
)

type myhandler struct{}

func (m *myhandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "TEST!")
}

func main() {
	mh := myhandler{}
	server := http.Server{
		Addr:    ":8080",
		Handler: &mh,
	}
	server.ListenAndServe()
}
