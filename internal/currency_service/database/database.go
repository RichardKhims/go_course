package database

import (
	"context"

	"github.com/RichardKhims/go_course/internal/currency_service/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Database interface {
	GetCourse(ctx context.Context, currency1 string, currency2 string) (result []Course, err error)
	CreateCourse(ctx context.Context, c Course) error
	DeleteCourse(ctx context.Context, c Course) error
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

type Course struct {
	tableName  struct{} `sql:"course"`
	ID         int
	Currency1  string
	Currency2  string
	Mean	   float64
}

func (d *DB) GetCourse(ctx context.Context, currency1 string, currency2 string) (result []Course, err error) {
	q := "SELECT id, currency1, currency2, mean FROM course WHERE currency1 = $1 and currency2 = $2;"
	if err = d.conn.SelectContext(ctx, &result, q, currency1, currency2); err != nil {
		return nil, err
	}
	return result, err
}

func (d *DB) CreateCourse(ctx context.Context, c Course) error {
	q := "INSERT INTO course (cur1, cur2) VALUES ($1, $2);"
	_, err := d.conn.ExecContext(ctx, q, c.Currency1, c.Currency2)
	return err
}

func (d *DB) DeleteCourse(ctx context.Context, c Course) error {
	q := "DELETE FROM course WHERE cur1 = $1 and cur2 = $2;"
	_, err := d.conn.ExecContext(ctx, q, c.Currency1, c.Currency2)
	return err
}

func (d *DB) Close() {
	d.conn.Close()
}