package order

type OrderDto struct {
	Name  string  `json:name`
	Price float64 `json:price`
	Date  string  `json:date`
}
