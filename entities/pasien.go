package entities

type Pasien struct {
	Id          int64
	Title  		string `validate:"required" label:"Title", min=20"`
	Content     string `validate:"required", min=200"`
	Category     string `validate:"required", min=3"`
	Status 		string `validate:"required" label:"Status !"`
}
