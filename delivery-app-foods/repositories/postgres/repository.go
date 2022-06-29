package postgres

import (
	"context"
	"database/sql"
	"time"

	"github.com/SantiagoBedoya/delivery-app-foods/foods"
)

const (
	createFoodQuery     = "INSERT INTO foods (name, vendor_id, description, category, type, ready_time, price) VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING ID"
	getFoodsQuery       = "SELECT id, name, vendor_id, description, category, type, ready_time, price, rating FROM foods"
	getFoodByIDQuery    = "SELECT id, name, vendor_id, description, category, type, ready_time, price, rating FROM foods WHERE id = $1"
	updateFoodByIDQuery = "UPDATE foods SET name=$1, description=$2, category=$3, type=$4, ready_time=$5, price=$6 WHERE id = $7"
	deleteFoodByIDQuery = "DELETE FROM foods WHERE id = $1"
)

type postgresRepository struct {
	db      *sql.DB
	timeout time.Duration
}

// NewPostgresRepository creates and implements accounts.Repository
func NewPostgresRepository(db *sql.DB) foods.Repository {
	return &postgresRepository{db: db, timeout: 2 * time.Second}
}

func (r *postgresRepository) Create(data *foods.Food) (*foods.Food, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	stmt, err := r.db.Prepare(createFoodQuery)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	if err := stmt.QueryRowContext(
		ctx,
		data.Name,
		data.VendorID,
		data.Description,
		data.Category,
		data.Type,
		data.ReadyTime,
		data.Price,
	).Scan(&data.ID); err != nil {
		return nil, err
	}
	return data, nil
}
func (r *postgresRepository) UpdateByID(foodID string, data *foods.Food) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	if _, err := r.GetByID(foodID); err != nil {
		return err
	}
	stmt, err := r.db.Prepare(updateFoodByIDQuery)
	if err != nil {
		return err
	}
	if _, err := stmt.ExecContext(
		ctx,
		data.Name,
		data.Description,
		data.Category,
		data.Type,
		data.ReadyTime,
		data.Price,
		foodID,
	); err != nil {
		return err
	}
	return nil
}
func (r *postgresRepository) DeleteByID(foodID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	stmt, err := r.db.Prepare(deleteFoodByIDQuery)
	if err != nil {
		return err
	}
	if _, err := stmt.ExecContext(ctx, foodID); err != nil {
		return err
	}
	return nil
}
func (r *postgresRepository) GetAll() ([]foods.Food, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	stmt, err := r.db.Prepare(getFoodsQuery)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return nil, err
	}
	results := make([]foods.Food, 0)
	for rows.Next() {
		var result foods.Food
		if err := rows.Scan(
			&result.ID,
			&result.Name,
			&result.VendorID,
			&result.Description,
			&result.Category,
			&result.Type,
			&result.ReadyTime,
			&result.Price,
			&result.Rating,
		); err != nil {
			return nil, err
		}
		results = append(results, result)
	}
	return results, nil
}
func (r *postgresRepository) GetByID(foodID string) (*foods.Food, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	stmt, err := r.db.Prepare(getFoodByIDQuery)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var result foods.Food
	if err := stmt.QueryRowContext(ctx, foodID).Scan(
		&result.ID,
		&result.Name,
		&result.VendorID,
		&result.Description,
		&result.Category,
		&result.Type,
		&result.ReadyTime,
		&result.Price,
		&result.Rating,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, foods.ErrFoodNotFound
		}
		return nil, err
	}
	return &result, nil
}
