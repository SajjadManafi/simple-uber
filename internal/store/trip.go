package store

import (
	"context"

	"github.com/SajjadManafi/simple-uber/models"
)

// createTripQuery is the query to create a trip
const createTripQuery = `
INSERT INTO trips (
  origin,
  destination,
  rider_id,
  driver_id,
  amount,
  cab_id
) VALUES (
  $1, $2, $3, 0, $4, 0
) RETURNING id, origin, destination, rider_id, driver_id, start_time, end_time, status, amount, cab_id, driver_rating
`

// CreateTrip creates a trip in the database
func (q *PostgresStore) CreateTrip(ctx context.Context, arg models.CreateTripParams) (models.Trip, error) {
	row := q.db.QueryRowContext(ctx, createTripQuery,
		arg.Origin,
		arg.Destination,
		arg.RiderID,
		arg.Amount,
	)
	var i models.Trip
	err := row.Scan(
		&i.ID,
		&i.Origin,
		&i.Destination,
		&i.RiderID,
		&i.DriverID,
		&i.StartTime,
		&i.EndTime,
		&i.Status,
		&i.Amount,
		&i.CabID,
		&i.DriverRating,
	)
	return i, err
}

// getTripQuery is the query to get a trip
const getTripQuery = `
SELECT id, origin, destination, rider_id, driver_id, start_time, end_time, status, amount, cab_id, driver_rating FROM trips
WHERE id = $1 LIMIT 1
`

// GetTrip gets a trip from the database
func (q *PostgresStore) GetTrip(ctx context.Context, id int32) (models.Trip, error) {
	row := q.db.QueryRowContext(ctx, getTripQuery, id)
	var i models.Trip
	err := row.Scan(
		&i.ID,
		&i.Origin,
		&i.Destination,
		&i.RiderID,
		&i.DriverID,
		&i.StartTime,
		&i.EndTime,
		&i.Status,
		&i.Amount,
		&i.CabID,
		&i.DriverRating,
	)
	return i, err
}

// listTripsQuery is the query to list trips
const listTripsQuery = `
SELECT id, origin, destination, rider_id, driver_id, start_time, end_time, status, amount, cab_id, driver_rating FROM trips
WHERE 
    driver_id = $1 OR rider_id = $2
ORDER BY id
LIMIT $3
OFFSET $4
`

// ListTrips lists trips from the database with driver or rider id
func (q *PostgresStore) ListTrips(ctx context.Context, arg models.ListTripsParams) ([]models.Trip, error) {
	rows, err := q.db.QueryContext(ctx, listTripsQuery,
		arg.DriverID,
		arg.RiderID,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []models.Trip{}
	for rows.Next() {
		var i models.Trip
		if err := rows.Scan(
			&i.ID,
			&i.Origin,
			&i.Destination,
			&i.RiderID,
			&i.DriverID,
			&i.StartTime,
			&i.EndTime,
			&i.Status,
			&i.Amount,
			&i.CabID,
			&i.DriverRating,
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

// updateTripStatusQuery is the query to update a trip status
const updateTripStatusQuery = `
UPDATE trips 
SET status = $2
WHERE id = $1
RETURNING id, origin, destination, rider_id, driver_id, start_time, end_time, status, amount, cab_id, driver_rating
`

// UpdateTripStatus updates a trip status
func (q *PostgresStore) UpdateTripStatus(ctx context.Context, arg models.UpdateTripStatusParams) (models.Trip, error) {
	row := q.db.QueryRowContext(ctx, updateTripStatusQuery, arg.ID, arg.Status)
	var i models.Trip
	err := row.Scan(
		&i.ID,
		&i.Origin,
		&i.Destination,
		&i.RiderID,
		&i.DriverID,
		&i.StartTime,
		&i.EndTime,
		&i.Status,
		&i.Amount,
		&i.CabID,
		&i.DriverRating,
	)
	return i, err
}

// updateTripDriverAccpetQuery is the query to update a trip to driver accept
const updateTripDriverAccpetQuery = `
UPDATE trips 
SET status = 2, driver_id = $2, cab_id = $3
WHERE id = $1
RETURNING id, origin, destination, rider_id, driver_id, start_time, end_time, status, amount, cab_id, driver_rating
`

// UpdateTripDriverAccpet updates a trip to driver accept
func (q *PostgresStore) UpdateTripDriverAccpet(ctx context.Context, arg models.UpdateTripDriverAccpetParams) (models.Trip, error) {
	row := q.db.QueryRowContext(ctx, updateTripDriverAccpetQuery, arg.ID, arg.DriverID, arg.CabID)
	var i models.Trip
	err := row.Scan(
		&i.ID,
		&i.Origin,
		&i.Destination,
		&i.RiderID,
		&i.DriverID,
		&i.StartTime,
		&i.EndTime,
		&i.Status,
		&i.Amount,
		&i.CabID,
		&i.DriverRating,
	)
	return i, err
}

// updateTripDoneQuery is the query to update a trip to done
const updateTripDoneQuery = `
UPDATE trips 
SET status = 4, end_time = now(), driver_rating = $2
WHERE id = $1
RETURNING id, origin, destination, rider_id, driver_id, start_time, end_time, status, amount, cab_id, driver_rating
`

// UpdateTripDone updates a trip to done
func (q *PostgresStore) UpdateTripDone(ctx context.Context, arg models.UpdateTripDoneParams) (models.Trip, error) {
	row := q.db.QueryRowContext(ctx, updateTripDoneQuery, arg.ID, arg.DriverRating)
	var i models.Trip
	err := row.Scan(
		&i.ID,
		&i.Origin,
		&i.Destination,
		&i.RiderID,
		&i.DriverID,
		&i.StartTime,
		&i.EndTime,
		&i.Status,
		&i.Amount,
		&i.CabID,
		&i.DriverRating,
	)
	return i, err
}

// deleteTripQuery is the query to delete a trip
const deleteTripQuery = `
DELETE FROM trips WHERE id = $1
`

// DeleteTrip deletes a trip from the database
func (q *PostgresStore) DeleteTrip(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteTripQuery, id)
	return err
}
