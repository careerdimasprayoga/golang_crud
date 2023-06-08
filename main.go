package main

import (
	"net/http"
	"github.com/jeypc/go-crud/controllers/pasiencontroller"
	"controllers/postcontroller"
)

func main() {
	// http.HandleFunc("/article", pasiencontroller.All_post)
	// http.HandleFunc("/article/", pasiencontroller.All_post)
	// http.HandleFunc("/article", pasiencontroller.All_post)

	http.HandleFunc("/all_post", postcontroller.All_post)

	http.HandleFunc("/", pasiencontroller.Index)
	http.HandleFunc("/pasien", pasiencontroller.Index)
	http.HandleFunc("/pasien/index", pasiencontroller.Index)
	http.HandleFunc("/pasien/add", pasiencontroller.Add)
	http.HandleFunc("/pasien/edit", pasiencontroller.Edit)
	http.HandleFunc("/pasien/delete", pasiencontroller.Delete)

	http.ListenAndServe(":3000", nil)
}
