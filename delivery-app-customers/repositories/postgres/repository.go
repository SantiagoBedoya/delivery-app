package postgres

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/SantiagoBedoya/delivery-app-customers/accounts"
	"github.com/jackc/pgconn"
)

const (
	createCustomerQuery      = "INSERT INTO customers(first_name, last_name, email, password, address, phone) VALUES ($1, $2, $3, $4, $5, $6) RETURNING ID"
	findCustomerByEmailQuery = "SELECT id, first_name, last_name, email, password FROM customers WHERE email = $1"
)

type postgresRepository struct {
	db      *sql.DB
	timeout time.Duration
}

// NewPostgresRepository creates and implements accounts.Repository
func NewPostgresRepository(db *sql.DB) accounts.Repository {
	return &postgresRepository{db: db, timeout: 2 * time.Second}
}

func (r *postgresRepository) Create(data *accounts.Customer) (*accounts.Customer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	stmt, err := r.db.Prepare(createCustomerQuery)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var insertedID int
	if err := stmt.QueryRowContext(
		ctx,
		data.FirstName,
		data.LastName,
		data.Email,
		data.Password,
		data.Address,
		data.Phone,
	).Scan(&insertedID); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.SQLState() == "23505" {
			return nil, accounts.ErrDuplicateEmail
		}
		return nil, err
	}
	data.ID = insertedID
	return data, nil
}

func (r *postgresRepository) FindByEmail(email string) (*accounts.Customer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	stmt, err := r.db.Prepare(findCustomerByEmailQuery)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var result accounts.Customer
	if err := stmt.QueryRowContext(ctx, email).Scan(
		&result.ID,
		&result.FirstName,
		&result.LastName,
		&result.Email,
		&result.Password,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, accounts.ErrAccountNotFound
		}
		return nil, err
	}
	return &result, nil
}
