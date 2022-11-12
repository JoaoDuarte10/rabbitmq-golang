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

type OrderRepositorySqlite struct {
	Db *sql.DB
}

func (o *OrderRepositorySqlite) Save(order order.OrderDto) error {
	stmt, err := o.Db.Prepare("INSERT INTO orders(name, price, date) VALUES(?, ?, ?)")
	if err != nil {
		return err
	}

	if _, err := stmt.Exec(order.Name, order.Price, order.Date); err != nil {
		return err
	}

	defer o.Db.Close()
	return nil
}

func (o *OrderRepositorySqlite) Get() ([]order.OrderDto, error) {
	result := []order.OrderDto{}

	rows, err := o.Db.Query("SELECT * FROM orders")
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

	defer o.Db.Close()

	return result, nil
}
