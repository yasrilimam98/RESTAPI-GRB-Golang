package models

import (
	"net/http"

	"github.com/yasrilimam98/grb-restapi/database"
)

type Client struct {
	// Id     int    `json:"id"`
	Regno  int    `json:"regno"`
	Serial string `json:"serial"`
}

func FetchAllClient(regno int, serial string) (Response, error) {
	var obj Client
	var arrobj []Client
	var res Response
	con := database.CreateCondb2()

	// sqlStatement := "SELECT CASE WHEN DATEDIFF(duedate,curdate()) < 0 THEN 0 ELSE 1 END as statuslisensi, DATEDIFF(duedate,curdate()) as jmlhari,dbname,dbport,localnetwork,localport,publicnetwork,publicport,`name` FROM mclient Where regno = ? AND serial = ?",regno,serial)

	rows, err := con.Query("SELECT CASE WHEN DATEDIFF(duedate,curdate()) < 0 THEN 0 ELSE 1 END as statuslisensi, DATEDIFF(duedate,curdate()) as jmlhari,dbname,dbport,localnetwork,localport,publicnetwork,publicport,`name` FROM mclient Where regno = ? AND serial = ?",regno,serial)
	defer rows.Close()

	if err != nil {
		return res, err
	}
	for rows.Next() {
		err = rows.Scan(&obj.Regno, &obj.Serial)
		if err != nil {
			return res, err
		}
		arrobj = append(arrobj, obj)
	}
	res.Status = http.StatusOK
	res.Message = "Succes"
	res.Data = arrobj
	return res, nil
}
