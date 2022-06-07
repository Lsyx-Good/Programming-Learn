/*
 * @Author: your name
 * @Date: 2021-08-02 16:17:44
 * @LastEditTime: 2021-08-08 11:16:21
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \project\Handler\fileserver_handler.go
 */
package main

import "net/http"

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		http.ServeFile(rw, r, "web"+r.URL.Path)
	})
	// http.ListenAndServe(":8080", http.FileServer(http.Dir("web")))
	http.ListenAndServe(":8080", nil)
}
