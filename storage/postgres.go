package storage

import (
	"context"

	// github
	"github.com/jackc/pgx/v5/pgxpool"
)

type Storage struct {
	pool *pgxpool.Pool
}


func New(pool *pgxpool.Pool) *Storage {
	return &Storage{
		pool: pool,
	}
}

func (s *Storage) SaveUrl(ctx context.Context, originalUrl, hash string) error {
	respone := `INSERT INTO urls (original_url, hash) VALUES ($1, $2)`
	_, err := s.pool.Exec(ctx,respone,originalUrl,hash)
	return err
}

func (s *Storage) GetUrl(ctx context.Context, hash string) (string, error) {
	var originalUrl string
	respone := `SELECT original_url FROM urls WHERE hash = $1`
	err := s.pool.QueryRow(ctx,respone,hash).Scan(&originalUrl)
	return originalUrl,err
}