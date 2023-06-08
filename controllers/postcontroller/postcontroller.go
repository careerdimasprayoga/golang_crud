package postcontroller

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"github.com/careerdimasprayoga/golang_crud/libraries"
	"github.com/careerdimasprayoga/golang_crud/models"
	"github.com/careerdimasprayoga/golang_crud/entities"
)

var validation = libraries.NewValidation()
var postModel = models.NewPostModel()


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
		vErrors := validation.Struct(post)

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

func All_post(response http.ResponseWriter, request *http.Request) {
	postModel := models.NewPostModel() // Inisialisasi postModel

	page, err := strconv.Atoi(request.URL.Query().Get("page"))
	if err != nil {
		page = 1
	}

	perPage := 10
	offset := (page - 1) * perPage

	posts := postModel.GetPaginatedPosts(offset, perPage)
	totalPosts := postModel.CountPosts()

	data := struct {
		Posts       []entities.Post
		TotalPosts  int
		CurrentPage int
	}{
		Posts:       posts,
		TotalPosts:  totalPosts,
		CurrentPage: page,
	}

	temp, err := template.ParseFiles("views/post/all_post.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(response, data)
}
