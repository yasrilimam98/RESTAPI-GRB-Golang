package models

import (
	"net/http"

	"github.com/yasrilimam98/grb-restapi/database"
)

type Karyawan struct {
	Id      int    `json:"id"`
	Nama    string `json:"nama`
	Alamat  string `json:"alamat"`
	Telepon string `json:"telepon"`
}

func FetchAllKaryawan() (Response, error) {
	var obj Karyawan
	var arrobj []Karyawan
	var res Response

	con := database.CreateCon()

	sqlStatement := "SELECT * FROM dt_karyawan"

	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	if err != nil {
		return res, err
	}
	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.Nama, &obj.Alamat, &obj.Telepon)
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

func StoreKaryawan(nama string, alamat string, telepon string) (Response, error) {
	var res Response

	// v := validator.New()

	// peg := Karyawan{
	// 	Nama:    nama,
	// 	Alamat:  alamat,
	// 	Telepon: telepon,
	// }

	// err := v.Struct(peg)
	// if err != nil {
	// 	return res, err
	// }

	con := database.CreateCon()

	sqlStatement := "INSERT dt_karyawan (nama, alamat, telepon) VALUES (?, ?, ?)"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(nama, alamat, telepon)
	if err != nil {
		return res, err
	}

	lastInsertedId, err := result.LastInsertId()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"last_inserted_id": lastInsertedId,
	}

	return res, nil
}

func UpdateKaryawan(id int, nama string, alamat string, telepon string) (Response, error) {
	var res Response

	con := database.CreateCon()

	sqlStatement := "UPDATE dt_karyawan SET nama = ?, alamat = ?, telepon = ? WHERE id = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(nama, alamat, telepon, id)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil
}

func DeleteKaryawan(id int) (Response, error) {
	var res Response

	con := database.CreateCon()

	sqlStatement := "DELETE FROM dt_karyawan WHERE id = ?"

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(id)
	if err != nil {
		return res, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"rows_affected": rowsAffected,
	}

	return res, nil
}
