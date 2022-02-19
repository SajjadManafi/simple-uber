package store

import (
	"context"

	"github.com/SajjadManafi/simple-uber/models"
)

// createCabQuery is a query for creating a cab.
const createCabQuery = `
INSERT INTO cabs (
  driver_id,
  brand,
  model,
  color,
  plate
) VALUES (
  $1, $2, $3, $4, $5
) RETURNING id, driver_id, brand, model, color, plate, created_at
`

// CreateCab creates a cab in the database.
func (q *PostgresStore) CreateCab(ctx context.Context, arg models.CreateCabParams) (models.Cab, error) {
	row := q.db.QueryRowContext(ctx, createCabQuery,
		arg.DriverID,
		arg.Brand,
		arg.Model,
		arg.Color,
		arg.Plate,
	)
	var i models.Cab
	err := row.Scan(
		&i.ID,
		&i.DriverID,
		&i.Brand,
		&i.Model,
		&i.Color,
		&i.Plate,
		&i.CreatedAt,
	)
	return i, err
}

// getCabQuery is a query for getting a cab.
const getCabQuery = `
SELECT id, driver_id, brand, model, color, plate, created_at FROM cabs
WHERE id = $1 LIMIT 1
`

// GetCab gets a cab from the database.
func (q *PostgresStore) GetCab(ctx context.Context, id int32) (models.Cab, error) {
	row := q.db.QueryRowContext(ctx, getCabQuery, id)
	var i models.Cab
	err := row.Scan(
		&i.ID,
		&i.DriverID,
		&i.Brand,
		&i.Model,
		&i.Color,
		&i.Plate,
		&i.CreatedAt,
	)
	return i, err
}

// listCabsQuery is a query that returns a list of cabs.
const listCabsQuery = `
SELECT id, driver_id, brand, model, color, plate, created_at FROM cabs
WHERE driver_id = $1
ORDER BY id
LIMIT $2
OFFSET $3
`

// ListCabs returns a list of cabs from database.
func (q *PostgresStore) ListCabs(ctx context.Context, arg models.ListCabsParams) ([]models.Cab, error) {
	rows, err := q.db.QueryContext(ctx, listCabsQuery, arg.DriverID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []models.Cab{}
	for rows.Next() {
		var i models.Cab
		if err := rows.Scan(
			&i.ID,
			&i.DriverID,
			&i.Brand,
			&i.Model,
			&i.Color,
			&i.Plate,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

// updateCabsDriverIdQuery is a query that updates the driver_id of a cab.
const updateCabsDriverIdQuery = `
UPDATE cabs 
SET driver_id = $2
WHERE id = $1
RETURNING id, driver_id, brand, model, color, plate, created_at
`

// UpdateCabsDriverId updates the driver_id of a cab.
func (q *PostgresStore) UpdateCabsDriverId(ctx context.Context, arg models.UpdateCabsDriverIdParams) (models.Cab, error) {
	row := q.db.QueryRowContext(ctx, updateCabsDriverIdQuery, arg.ID, arg.DriverID)
	var i models.Cab
	err := row.Scan(
		&i.ID,
		&i.DriverID,
		&i.Brand,
		&i.Model,
		&i.Color,
		&i.Plate,
		&i.CreatedAt,
	)
	return i, err
}

// deleteCabQuery is a query for deleting a cab.
const deleteCabQuery = `
DELETE FROM cabs WHERE id = $1
`

// DeleteCab deletes a cab from the database.
func (q *PostgresStore) DeleteCab(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteCabQuery, id)
	return err
}
