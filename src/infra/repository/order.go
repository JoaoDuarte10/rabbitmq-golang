package repository

import (
	"database/sql"
	"log"
	"rabbitmq-golang/src/domain/order"
)

type OrderRepository interface {
	Save(order order.OrderDto) error
	Get() ([]order.OrderDto, error)
}

type OrderRepositorySqlite struct{}

func (o *OrderRepositorySqlite) Save(order order.OrderDto) error {
	db, err := sql.Open("sqlite3", "./order.db")
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO orders(name, price, date) VALUES(?, ?, ?)")
	if err != nil {
		return err
	}

	if _, err := stmt.Exec(order.Name, order.Price, order.Date); err != nil {
		return err
	}
	return nil
}

func (o *OrderRepositorySqlite) Get() ([]order.OrderDto, error) {
	db, err := sql.Open("sqlite3", "./order.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	result := []order.OrderDto{}

	rows, err := db.Query("SELECT * FROM orders")
	if err != nil {
		log.Printf("Failed to get orders: %s", err)
		return nil, err
	}

	for rows.Next() {
		var order order.OrderDto
		if err := rows.Scan(&order.Id, &order.Name, &order.Price, &order.Date); err != nil {
			log.Printf("Failed to parsed orders: %s", err)
		}
		result = append(result, order)
	}

	return result, nil
}
