package store

import (
	"L0/internal/caching"
	"L0/internal/model"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

// Store is a database structure
type Store struct {
	dbpool *pgxpool.Pool
}

// NewStore initialized service store
func NewStore(url string) (*Store, error) {
	var err error
	dbpool, err := pgxpool.Connect(context.Background(), url)
	if err != nil {
		return &Store{}, err
	}

	return &Store{dbpool: dbpool}, nil
}

// AddOrder insert order into database
func (s *Store) AddOrder(order model.Order) error {
	_, err := s.dbpool.Exec(context.Background(), "insert into orders (\"order\") values ($1)", order)
	if err != nil {
		return err
	}
	return nil
}

// getOrder return order from database
func (s *Store) getOrder(id string) (model.Order, error) {
	var order model.Order
	err := s.dbpool.QueryRow(context.Background(), "select \"order\" from orders where id=$1", id).Scan(&order)
	if err != nil {
		return model.Order{}, err
	}

	return order, nil
}

// RestoreCache fill service cache from database
func (s *Store) RestoreCache(cache *caching.Cache) error {
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
