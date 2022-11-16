package database

import (
	"database/sql"

	"github.com/thenicolauuu/go-rabbitmq/internal/order/entity"
)

type OrderRepository struct {
	DB *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{DB: db}
}

func (r *OrderRepository) Save(order *entity.Order) error {
	stmt, err := r.DB.Prepare("INSERT INTO orders (id, price, tax, final_price) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(order.ID, order.Price, order.Price, order.FinalPrice)
	if err != nil {
		return err
	}

	return nil

}

func (r *OrderRepository) GetTotal() (int, error) {
	var total int
	err := r.DB.QueryRow("Select count(*) from orders").Scan(&total)
	if err != nil {
		return 0, err
	}

	return total, nil
}
