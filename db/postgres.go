package db

import "database/sql"

func New(driver, dataSourceName string) (*sql.DB, error) {
	db, err := sql.Open(driver, dataSourceName)
	if err != nil {
		return nil, err
	}
	return db, nil
}
