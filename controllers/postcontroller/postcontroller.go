package postcontroller

import (
	"html/template"
	"net/http"
	// "strconv"
	// "github.com/careerdimasprayoga/golang_crud/libraries"
	// "github.com/careerdimasprayoga/golang_crud/models"
	// "github.com/careerdimasprayoga/golang_crud/entities"
)

func all_post(response http.ResponseWriter, request *http.Request) {
	temp, err := template.ParseFiles("views/post/all_post.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(response, nil)
}

func add_post(response http.ResponseWriter, request *http.Request) {

	if request.Method == http.MethodGet {
		temp, err := template.ParseFiles("views/post/add_post.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(response, nil)
	}

}