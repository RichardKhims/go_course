package database

import (
	"context"

	"github.com/RichardKhims/go_course/internal/currency_service/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Database interface {
	GetAllCourses(ctx context.Context) (result []Course, err error)
	GetCourse(ctx context.Context, currency1 string, currency2 string) (result []Course, err error)
	CreateCourse(ctx context.Context, c Course) error
	DeleteCourse(ctx context.Context, c Course) error
	UpdateCourse(ctx context.Context, currency1 string, currency2 string, mean float64, last_changed string) error
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

func (d *DB) GetAllCourses(ctx context.Context) (result []Course, err error) {
	q := "SELECT id, currency1, currency2, mean FROM course;"
	if err = d.conn.SelectContext(ctx, &result, q); err != nil {
		return nil, err
	}
	return result, err
}

func (d *DB) GetCourse(ctx context.Context, currency1 string, currency2 string) (result []Course, err error) {
	q := "SELECT id, currency1, currency2, mean FROM course WHERE currency1 = $1 and currency2 = $2;"
	if err = d.conn.SelectContext(ctx, &result, q, currency1, currency2); err != nil {
		return nil, err
	}
	return result, err
}

func (d *DB) CreateCourse(ctx context.Context, c Course) error {
	q := "INSERT INTO course (currency1, currency2) VALUES ($1, $2);"
	_, err := d.conn.ExecContext(ctx, q, c.Currency1, c.Currency2)
	return err
}

func (d *DB) DeleteCourse(ctx context.Context, c Course) error {
	q := "DELETE FROM course WHERE currency1 = $1 and currency2 = $2;"
	_, err := d.conn.ExecContext(ctx, q, c.Currency1, c.Currency2)
	return err
}

func (d *DB) UpdateCourse(ctx context.Context, currency1 string, currency2 string, mean float64, last_changed string) error {
	q := "UPDATE course SET mean = $1, last_changed = $2 WHERE currency1 = $3 and currency2 = $4;"
	_, err := d.conn.ExecContext(ctx, q, mean, last_changed, currency1, currency2)
	return err
}

func (d *DB) Close() {
	d.conn.Close()
}