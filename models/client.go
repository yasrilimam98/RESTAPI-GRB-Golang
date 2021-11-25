package models

import (
	"github.com/yasrilimam98/grb-restapi/database"
)

func FetchAllClient(regno int, serial string) ([]Client, error) {
	con := database.CreateCondb2()
	rows, err := con.Query("SELECT CASE WHEN DATEDIFF(duedate,curdate()) < 0 THEN 0 ELSE 1 END , DATEDIFF(duedate,curdate()),dbname,dbport,localnetwork,localport,publicnetwork,publicport,name FROM mclient Where regno = ? AND serial = ?", regno, serial)

	if err != nil {
		return nil, err
	} else {
		clients := []Client{}
		for rows.Next() {
			// var regno int
			// var serial string
			// var duedate string
			// var name string
			// var dbname string
			// var dbport string
			// var localnetwork string
			// var localport string
			// var publicnetwork string
			// var publicport string
			var status1 string
			var sisa1 string
			var dbname1 string
			var dbport1 string
			var locnet1 string
			var locport1 string
			var pubnet1 string
			var pubport1 string
			var name1 string
			err2 := rows.Scan(&status1, &sisa1, &dbname1, &dbport1, &locnet1, &locport1, &pubnet1, &pubport1, &name1)
			if err2 != nil {
				return nil, err2
			} else {
				client := Client{regno, serial, status1, sisa1, dbname1, dbport1, locnet1, locport1, pubnet1, pubport1}
				clients = append(clients, client)
			}
		}
		return clients, nil
	}
}
