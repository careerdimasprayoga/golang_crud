package entities

type Post struct {
	Id          int64
	Title  		string `validate:"required,min_length=20" label:"Title""`
	Content     string `validate:"required,min_length=200"`
	Category     string `validate:"required,min_length=3"`
	Status 		string `validate:"required" label:"Status !"`
}
