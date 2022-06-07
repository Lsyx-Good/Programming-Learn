/*
 * @Author: your name
 * @Date: 2021-09-21 23:15:29
 * @LastEditTime: 2021-09-21 23:20:54
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \Go\src\project\Tempplate\tempplate.go
 */
package tempplate

import (
	"html/template"
	"net/http"
)

func Test(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("gh")
	t.Execute(w, "Hjksdhfjkdshf")
}

func main() {
	http.HandleFunc("/", Test)
	http.ListenAndServe(":8080", nil)
}
