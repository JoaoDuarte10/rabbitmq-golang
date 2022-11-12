package repository

import (
	"database/sql"
	"rabbitmq-golang/src/domain/order"
)

type OrderRepository interface {
	Save(order order.OrderDto) error
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
