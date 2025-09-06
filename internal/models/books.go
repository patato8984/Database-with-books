package models

type Books = struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Age         int    `json:"age"`
	Id_author   int    `json:"id_author"`
	Name_author string `json:"name_author"`
}
