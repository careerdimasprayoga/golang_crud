package entities

type Post struct {
	Id          int64
	Title  		string `validate:"required" label:"Title"`
	Content     string `validate:"required"`
	Category    string `validate:"required"`
	Status 		string `validate:"required" label:"Status !"`
}
