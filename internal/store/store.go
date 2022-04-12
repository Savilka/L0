package store

import (
	"L0/internal/caching"
	"L0/internal/model"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Store struct {
	dbpool *pgxpool.Pool
}

func NewStore(url string) (*Store, error) {
	var err error
	dbpool, err := pgxpool.Connect(context.Background(), url)
	if err != nil {
		return &Store{}, err
	}

	return &Store{dbpool: dbpool}, nil
}

func (s *Store) AddOrder(order model.Order) error {
	_, err := s.dbpool.Exec(context.Background(), "insert into orders (\"order\") values ($1)", order)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) GetOrder(id string) (model.Order, error) {
	var order model.Order
	err := s.dbpool.QueryRow(context.Background(), "select \"order\" from orders where id=$1", id).Scan(&order)
	if err != nil {
		return model.Order{}, err
	}

	return order, nil
}

func (s *Store) RestoreCash(cache *caching.Cache) error {
	rows, err := s.dbpool.Query(context.Background(), "select \"order\" from orders")
	if err != nil {
		return err
	}

	for rows.Next() {
		var order model.Order
		err = rows.Scan(&order)
		if err != nil {
			return err
		}

		cache.Put(order)
	}

	return rows.Err()
}
