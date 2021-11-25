package models

// import (
// 	"net/http"

// 	"github.com/yasrilimam98/grb-restapi/database"
// )

// type Client struct {
// 	Id            int    `json:"id"`
// 	Regno         int    `json:"regno"`
// 	Serial        string `json:"serial"`
// 	Date          string `json:"duedate"`
// 	Name          string `json:"name"`
// 	DBname        string `json:"dbname"`
// 	DBport        string `json:"dbport"`
// 	LocalNetwork  string `json:"localnetwork"`
// 	LocalPort     string `json:"localport"`
// 	PublicNetwork string `json:"publicnetwork"`
// 	PublicPort    string `json:"publicport"`
// }

// func FetchAllClient(regno int, serial string) (Response, error) {
// 	var obj Client
// 	var arrobj []Client
// 	var res Response
// 	con := database.CreateCondb2()

// 	// sqlStatement := "SELECT CASE WHEN DATEDIFF(duedate,curdate()) < 0 THEN 0 ELSE 1 END as statuslisensi, DATEDIFF(duedate,curdate()) as jmlhari,dbname,dbport,localnetwork,localport,publicnetwork,publicport,`name` FROM mclient Where regno = ? AND serial = ?",regno,serial)
// 	// rows, err := con.Query("SELECT CASE WHEN DATEDIFF(duedate,curdate()) < 0 THEN 0 ELSE 1 END as id, DATEDIFF(duedate,curdate()) as serial,dbname,dbport,localnetwork,localport,publicnetwork,publicport,`name` FROM mclient Where regno = ? AND serial = ?", regno, serial)
// 	rows, err := con.Query("call getclientduedate(?, ?)", regno, serial)
// 	// rows, err := con.Query("SELECT id, regno, serial, duedate, name, dbname, dbport,localnetwork, localport, publicnetwork, publicport FROM mclient Where regno = ? AND serial = ?", regno, serial)
// 	defer rows.Close()

// 	if err != nil {
// 		return res, err
// 	}
// 	for rows.Next() {
// 		err = rows.Scan(&obj.Id, &obj.Regno, &obj.Serial, &obj.Date, &obj.Name, &obj.DBname, &obj.DBport, &obj.LocalNetwork, &obj.LocalPort, &obj.PublicNetwork, &obj.PublicPort)
// 		if err != nil {
// 			return res, err
// 		}
// 		arrobj = append(arrobj, obj)
// 	}
// 	res.Status = http.StatusOK
// 	res.Message = "Succes"
// 	res.Data = arrobj
// 	return res,nil
// }
