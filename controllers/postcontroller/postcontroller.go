package postcontroller

import (
	"html/template"
	"net/http"
	// "strconv"
	"github.com/careerdimasprayoga/golang_crud/libraries"
	"github.com/careerdimasprayoga/golang_crud/models"
	"github.com/careerdimasprayoga/golang_crud/entities"
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
	} else if request.Method == http.MethodPost {
		request.ParseForm()
		var pasien entities.Pasien
		pasien.NamaLengkap = request.Form.Get("nama_lengkap")
		pasien.NIK = request.Form.Get("nik")
		pasien.JenisKelamin = request.Form.Get("jenis_kelamin")
		pasien.TempatLahir = request.Form.Get("tempat_lahir")
		pasien.TanggalLahir = request.Form.Get("tanggal_lahir")
		pasien.Alamat = request.Form.Get("alamat")
		pasien.NoHp = request.Form.Get("no_hp")

		var data = make(map[string]interface{})
		vErrors := validation.Struct(pasien)

		if vErrors != nil {
			data["pasien"] = pasien
			data["validation"] = vErrors
		} else {
			data["pesan"] = "Data pasien berhasil disimpan"
			// pasienModel.Create(pasien)
		}

		temp, _ := template.ParseFiles("views/pasien/add.html")
		temp.Execute(response, data)
	}

}