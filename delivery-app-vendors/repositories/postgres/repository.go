package postgres

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/SantiagoBedoya/delivery-app-vendors/vendors"
	"github.com/jackc/pgconn"
)

const (
	createVendorQuery      = "INSERT INTO vendors(name, owner_name, address, email, phone, password) VALUES ($1, $2, $3, $4, $5, $6) RETURNING ID"
	findVendorByEmailQuery = "SELECT id, name, owner_name, email, password FROM vendors WHERE email = $1"
)

type postgresRepository struct {
	db      *sql.DB
	timeout time.Duration
}

// NewPostgresRepository creates and implements accounts.Repository
func NewPostgresRepository(db *sql.DB) vendors.Repository {
	return &postgresRepository{db: db, timeout: 2 * time.Second}
}

func (r *postgresRepository) Create(data *vendors.Vendor) (*vendors.Vendor, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	stmt, err := r.db.Prepare(createVendorQuery)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var insertedID int
	if err := stmt.QueryRowContext(
		ctx,
		data.Name,
		data.OwnerName,
		data.Address,
		data.Email,
		data.Phone,
		data.Password,
	).Scan(&insertedID); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.SQLState() == "23505" {
			return nil, vendors.ErrDuplicateEmail
		}
		return nil, err
	}
	data.ID = insertedID
	return data, nil
}

func (r *postgresRepository) FindByEmail(email string) (*vendors.Vendor, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	stmt, err := r.db.Prepare(findVendorByEmailQuery)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var result vendors.Vendor
	if err := stmt.QueryRowContext(ctx, email).Scan(
		&result.ID,
		&result.Name,
		&result.OwnerName,
		&result.Email,
		&result.Password,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, vendors.ErrAccountNotFound
		}
		return nil, err
	}
	return &result, nil
}
