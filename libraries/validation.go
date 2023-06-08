package postcontroller

import (
	"html/template"
	"net/http"

	"github.com/careerdimasprayoga/golang_crud/libraries"
	"github.com/careerdimasprayoga/golang_crud/models"
	"github.com/careerdimasprayoga/golang_crud/entities"
)

var validation = libraries.NewValidation()
var postModel = models.NewPostModel()

func All_post(response http.ResponseWriter, request *http.Request) {
	temp, err := template.ParseFiles("views/post/all_post.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(response, nil)
}

func Add_post(response http.ResponseWriter, request *http.Request) {

	if request.Method == http.MethodGet {
		temp, err := template.ParseFiles("views/post/add_post.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(response, nil)
	} else if request.Method == http.MethodPost {
		request.ParseForm()
		var post entities.Post
		post.Title = request.Form.Get("title")
		post.Content = request.Form.Get("content")
		post.Category = request.Form.Get("category")
		post.Status = request.Form.Get("status")
		
		var data = make(map[string]interface{})
		_, vErrors := validation.Struct(post)

		if vErrors != nil {
			data["post"] = post
			data["validation"] = vErrors
		} else {
			data["pesan"] = "Post successfully saved!"
			postModel.Create(post)
		}

		temp, _ := template.ParseFiles("views/post/add_post.html")
		temp.Execute(response, data)
	}

}