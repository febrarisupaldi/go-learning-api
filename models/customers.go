package models

import (
	"net/http"
	"time"

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

func AddCustomers(customer_name string, customer_address string, customer_contact string) (Response, error) {
	var res Response
	con := db.Conn()
	time := time.Now()
	sqlStatement := "insert into customers(customer_name, customer_address, customer_contact, created_at, sales_id) values(?,?,?,?,120)"
	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(customer_name, customer_address, customer_contact, time)

	if err != nil {
		return res, err
	}

	getId, err := result.LastInsertId()
	if err != nil {
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{"id": getId}

	return res, nil
}
