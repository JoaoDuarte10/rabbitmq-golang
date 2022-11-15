package repository

import (
	"database/sql"
	"log"
	"rabbitmq-golang/src/domain/entity"
)

type OrderRepository interface {
	Save(order entity.OrderDto) error
	Get() ([]entity.OrderDto, error)
}

type OrderRepositorySqlite struct {
	Db *sql.DB
}

func (o *OrderRepositorySqlite) Save(order entity.OrderDto) error {
	stmt, err := o.Db.Prepare("INSERT INTO orders(name, price, date) VALUES(?, ?, ?)")
	if err != nil {
		return err
	}

	if _, err := stmt.Exec(order.Name, order.Price, order.Date); err != nil {
		return err
	}

	return nil
}

func (o *OrderRepositorySqlite) Get() ([]entity.OrderDto, error) {
	result := []entity.OrderDto{}

	rows, err := o.Db.Query("SELECT * FROM orders")
	if err != nil {
		log.Printf("Failed to get orders: %s", err)
		return nil, err
	}

	for rows.Next() {
		var order entity.OrderDto
		if err := rows.Scan(&order.Id, &order.Name, &order.Price, &order.Date); err != nil {
			log.Printf("Failed to parsed orders: %s", err)
		}
		result = append(result, order)
	}

	return result, nil
}
