package entities

type post struct {
	Id          int64
	title  		string `validate:"required" label:"Title"`
	content     string `validate:"required"`
	status 		string `validate:"required" label:"Status !"`
}
