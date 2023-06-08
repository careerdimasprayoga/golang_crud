package main

import (
	"net/http"
	// "github.com/jeypc/go-crud/controllers/pasiencontroller"
	"github.com/careerdimasprayoga/golang_crud/controllers/postcontroller"
)

func main() {
	// http.HandleFunc("/article", pasiencontroller.All_post)
	// http.HandleFunc("/article/", pasiencontroller.All_post)
	// http.HandleFunc("/article", pasiencontroller.All_post)

	http.HandleFunc("/", postcontroller.All_post)
	http.HandleFunc("/add_post", postcontroller.Add_post)
	// http.HandleFunc("/all_post", postcontroller.all_post)
	// http.HandleFunc("/drafts_post", postcontroller.drafts_post)
	// http.HandleFunc("/trashed_post", postcontroller.trashed_post)
	// http.HandleFunc("/edit_post", postcontroller.edit_post)
	// http.HandleFunc("/move_post", postcontroller.move_trashed_post)

	// http.HandleFunc("/", pasiencontroller.Index)
	// http.HandleFunc("/pasien", pasiencontroller.Index)
	// http.HandleFunc("/pasien/index", pasiencontroller.Index)
	// http.HandleFunc("/pasien/add", pasiencontroller.Add)
	// http.HandleFunc("/pasien/edit", pasiencontroller.Edit)
	// http.HandleFunc("/pasien/delete", pasiencontroller.Delete)

	http.ListenAndServe(":3000", nil)
}
