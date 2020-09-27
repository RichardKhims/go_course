package database

import (
	"context"

	"github.com/RichardKhims/go_course/internal/currency_service/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Database interface {
	GetCourse(ctx context.Context, cur1 string, cur2 string) (result []Course, err error)
	CreateCurrency(ctx context.Context, currency string, name string) error
	DeleteCurrency(ctx context.Context, currency string) error
	Close()
}

type DB struct {
	conn *sqlx.DB
}

func New(cfg config.DatabaseConfig) (*DB, error) {
	conn, err := sqlx.Connect("postgres", cfg.ConnectionString)
	if err != nil {
		return nil, err
	}
	return &DB{
		conn: conn,
	}, nil
}

type Currency struct {
	tableName struct{} `sql:"currency"`
	ID        int
	Symbol	  string
	Name	  string
}

type Course struct {
	tableName  struct{} `sql:"course"`
	ID         int
	Currency1  string
	Currency2  string
	mean	   float64
}

func (d *DB) GetCourse(ctx context.Context, cur1 string, cur2 string) (result []Course, err error) {
	q := "SELECT id, cur1, cur2, mean FROM course WHERE cur1 = $1 and cur2 = $2;"
	if err = d.conn.SelectContext(ctx, &result, q, cur1, cur2); err != nil {
		return nil, err
	}
	return result, err
}

func (d *DB) CreateCurrency(ctx context.Context, currency string, name string) error {
	q := "INSERT INTO currency (symbol, curname) VALUES ($1, $2);"
	_, err := d.conn.ExecContext(ctx, q, currency, name)
	return err
}

func (d *DB) DeleteCurrency(ctx context.Context, currency string) error {
	q := "DELETE FROM currency WHERE symbol = $1;"
	_, err := d.conn.ExecContext(ctx, q, currency)
	return err
}

func (d *DB) Close() {
	d.conn.Close()
}