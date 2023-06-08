package postcontroller

import (
	"html/template"
	"net/http"
	"strconv"
)

func All_post(response http.ResponseWriter, request *http.Request) {
	temp, err := template.ParseFiles("views/post/all_post.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(response, nil)
}
