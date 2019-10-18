package main

import (
	"log"
	"time"

	"github.com/samtech09/train-api/models"

	"github.com/jackc/pgx"
)

var (
	_pdb *pgx.ConnPool
)

//InitDB Initialize database connection for PostgreSQL database
func initDB() {
	var err error
	pgxConfig := pgx.ConnConfig{
		Host:     "192.168.60.206",
		Port:     5432,
		Database: "testdb",
		User:     "testuser",
		Password: "testuser",
	}
	pgxConnPoolConfig := pgx.ConnPoolConfig{pgxConfig, 8, nil, 5 * time.Second}
	_pdb, err = pgx.NewConnPool(pgxConnPoolConfig)
	if err != nil {
		log.Fatalf("Psql Connection error %v\n", err)
	}
}

func getRecordsFromDB() ([]models.Item, error) {
	sql := "select id, title, price from items;"

	rows, err := _pdb.Query(sql) //execute query and get records
	if err != nil {
		return nil, err
	}
	defer rows.Close() // make sure to always close rows

	var list []models.Item
	for rows.Next() {
		u := models.Item{} // scan rows to struct
		err := rows.Scan(&u.ID, &u.Title, &u.Price)
		if err != nil {
			return nil, err
		}
		list = append(list, u)
	}

	return list, nil
}

func getByIDFromDB(id int) (models.Item, error) {
	sql := "select id, title, price from items where id=$1;"

	u := models.Item{}
	rows, err := _pdb.Query(sql, id) //execute query and get records
	if err != nil {
		return u, err
	}
	defer rows.Close() // make sure to always close rows

	if rows.Next() {
		err := rows.Scan(&u.ID, &u.Title, &u.Price) // scan rows to struct
		if err != nil {
			return u, err
		}
	}
	return u, nil
}

func saveItem(itm models.Item) error {
	sql := "insert into items(Title, Price) values($1, $2);"

	_, err := _pdb.Exec(sql, itm.Title, itm.Price)
	if err != nil {
		return err
	}
	return nil
}

// func getRecordsFromDB() ([]student, error) {
// 	sql := "select roll, name, count(*) over as cnt from students"

// 	//get list of usres not submitted given test
// 	rows, err := _pdb.Query(sql)
// 	if err != nil {
// 		return nil, err
// 	}
// 	// make sure to always close rows
// 	defer rows.Close()

// 	var list []student
// 	count := 0
// 	i := 0
// 	for rows.Next() {
// 		// scan rows to struct
// 		u := student{}
// 		err := rows.Scan(&u.Roll, &u.Name, &count)
// 		if err != nil {
// 			return nil, err
// 		}

// 		if i == 0 {
// 			// instead of append, make slice of fixed length to avoid reallocations
// 			list = make([]student, count)
// 		}
// 		list[i] = u
// 		i++
// 	}

// 	return list, nil
// }
