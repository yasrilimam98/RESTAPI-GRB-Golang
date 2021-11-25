package models

type Response struct {
	Status int	`json:"status"`
	Message string	`json:"message"`
	Data interface{}	`json:"data"`
}

// type response struct {
// 	Status bool
// 	Pesan  string
// 	Data   []Client
// }
