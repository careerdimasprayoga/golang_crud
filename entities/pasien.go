package entities

type Pasien struct {
	Id          int64
	Title  		string `validate:"required,minstring=20" label:"Title""`
	Content     string `validate:"required,minstring=200"`
	Category     string `validate:"required,minstring=3"`
	Status 		string `validate:"required" label:"Status !"`
}
