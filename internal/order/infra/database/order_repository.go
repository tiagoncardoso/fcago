package database

import (
	"aceleracao/golang/internal/order/entity"
	"database/sql"
)

type OrderRepository struct {
	Db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{Db: db}
}

// sql create table on mysql
// CREATE TABLE `orders` (
//   `id` varchar(255) NOT NULL,
//   `price` float NOT NULL,
//   `tax` float NOT NULL,
//   `final_price` float NOT NULL,
//   PRIMARY KEY (`id`))
// )

func (r *OrderRepository) Save(order *entity.Order) error {
	stmt, err := r.Db.Prepare("INSERT INT orders (id, price, tax, final_price) VALUES (?, ?, ?, ?)")

	if err != nil {
		return err
	}

	_, err = stmt.Exec(order.ID, order.Price, order.Tax, order.FinalPrice)
	if err != nil {
		return err
	}

	return nil
}
