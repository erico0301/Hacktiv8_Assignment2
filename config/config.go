package config

import "database/sql"

func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/orders_by")
	if err != nil {
		return nil, err
	}
	return db, nil
}
