package database

import (
    "database/sql"

    "github.com/amichelins/goexpert_clean_arch/internal/entity"
)

type OrderRepository struct {
    Db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
    return &OrderRepository{Db: db}
}

func (r *OrderRepository) Save(order *entity.Order) error {

    stmt, err := r.Db.Prepare("INSERT INTO orders (id, price, tax, final_price) VALUES (?, ?, ?, ?)")

    if err != nil {
        return err
    }

    _, err = stmt.Exec(order.ID, order.Price, order.Tax, order.FinalPrice)

    if err != nil {
        return err
    }
    return nil
}

func (r *OrderRepository) List() (*[]entity.Order, error) {
    var sSql string
    var Orders []entity.Order
    var id string
    var price float64
    var tax float64
    var final_price float64

    sSql = " SELECT id, price, tax, final_price "
    sSql += " FROM orders "

    rows, err := r.Db.Query(sSql)

    if err != nil {
        return nil, err
    }

    for rows.Next() {

        err = rows.Scan(&id, &price, &tax, &final_price)

        if err != nil {
            return nil, err
        }
        Orders = append(Orders, entity.Order{ID: id, Price: price, Tax: tax, FinalPrice: final_price})
    }

    return &Orders, nil
}

func (r *OrderRepository) ListById(id string) (*entity.Order, error) {
    var sSql string
    var price float64
    var tax float64
    var final_price float64

    sSql = " SELECT price, tax, final_price "
    sSql += " FROM orders "
    sSql += " WHERE id = ? "

    err := r.Db.QueryRow(sSql, id).Scan(&price, &tax, &final_price)

    if err != nil {
        return nil, err

    }

    return &entity.Order{ID: id, Price: price, Tax: tax, FinalPrice: final_price}, nil
}
