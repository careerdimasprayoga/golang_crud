package entities

type Pasien struct {
	Id          int64
	Title  		string `validate:"required",min=20" label:"Title"`
	Content     string `validate:"required,min=200""`
	Category     string `validate:"required,min=3""`
	Status 		string `validate:"required" label:"Status !"`
}
