package models

// import (
// 	"net/http"

// 	"github.com/yasrilimam98/grb-restapi/database"
// )

// type Client struct {
// 	Id     int    `json:"id"`
// 	Regno  string `json:"regno"`
// 	Serial string `json:"serial"`
// }

// func FetchAllClient(regno string, serial string) (Response, error) {
// 	var obj Client
// 	var arrobj []Client
// 	var res Response
// 	con := database.CreateCondb2()

// 	sqlStatement := "SELECT a.regno, a.serial FROM mclient_test AS a "

// 	rows, err := con.Prepare(sqlStatement)

// 	if err != nil {
// 		return res, err
// 	}
// 	defer rows.Close()
// 	row := rows.QueryRow(regno, serial)
// 	if err := row.Scan(&obj.Regno, &obj.Serial); err != nil {
// 		return res, err
// 	}
// 	res.Status = http.StatusOK
// 	res.Message = "Succes"
// 	res.Data = arrobj
// 	return res, nil
// }
