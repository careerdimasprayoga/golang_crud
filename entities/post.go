package entities

type Post struct {
	Id          int64
	title  		string `validate:"required" label:"Title"`
	content     string `validate:"required"`
	category    string `validate:"required"`
	status 		string `validate:"required" label:"Status !"`
}
