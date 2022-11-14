package entity

type OrderDto struct {
	Id    any     `json:id`
	Name  string  `json:date`
	Price float64 `json:price`
	Date  string  `json:date`
}
