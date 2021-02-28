package models

import (
	"net/http"

	"github.com/febrarisupaldi/go-learning-api/db"
)

type Customer struct {
	Id      int    `json:"id"`
	Name    string `json:"customer_name"`
	Address string `json:"customer_address"`
	Contact string `json:"customer_contact"`
}

func GetAllCustomer() (Response, error) {
	var obj Customer
	var arrobj []Customer
	var res Response

	con := db.Conn()

	sqlStatement := "select id, customer_name, customer_address, customer_contact from customers"

	rows, err := con.Query(sqlStatement)
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.Name, &obj.Address, &obj.Contact)
		if err != nil {
			return res, err
		}

		arrobj = append(arrobj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "success"
	res.Data = arrobj

	return res, nil
}
