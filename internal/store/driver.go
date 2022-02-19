package store

import (
	"context"

	"github.com/SajjadManafi/simple-uber/models"
)

// createDriverQuery is query for insert driver
const createDriverQuery = `
INSERT INTO drivers (
  username,
  hashed_password,
  full_name,
  gender,
  balance,
  email,
  current_cab_id
) VALUES (
  $1, $2, $3, $4, 0, $5, 0
) RETURNING id, username, hashed_password, full_name, gender, balance, email, current_cab_id, joined_at
`

// CreateDriver creates driver in database
func (q *PostgresStore) CreateDriver(ctx context.Context, arg models.CreateDriverParams) (models.Driver, error) {
	row := q.db.QueryRowContext(ctx, createDriverQuery,
		arg.Username,
		arg.HashedPassword,
		arg.FullName,
		arg.Gender,
		arg.Email,
	)
	var i models.Driver
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.HashedPassword,
		&i.FullName,
		&i.Gender,
		&i.Balance,
		&i.Email,
		&i.CurrentCabID,
		&i.JoinedAt,
	)
	return i, err
}

// getDriverQuery is query for get driver by id
const getDriverQuery = `
SELECT id, username, hashed_password, full_name, gender, balance, email, current_cab_id, joined_at FROM drivers
WHERE id = $1 LIMIT 1
`

// GetDriver gets user from database by id
func (q *PostgresStore) GetDriver(ctx context.Context, id int32) (models.Driver, error) {
	row := q.db.QueryRowContext(ctx, getDriverQuery, id)
	var i models.Driver
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.HashedPassword,
		&i.FullName,
		&i.Gender,
		&i.Balance,
		&i.Email,
		&i.CurrentCabID,
		&i.JoinedAt,
	)
	return i, err
}

// listDriversQuery is query for get list of drivers from database
const listDriversQuery = `
SELECT id, username, hashed_password, full_name, gender, balance, email, current_cab_id, joined_at FROM drivers
ORDER BY id
LIMIT $1
OFFSET $2
`

// ListDrivers gets drivers and return slice of users (with limit and offset)
func (q *PostgresStore) ListDrivers(ctx context.Context, arg models.ListDriversParams) ([]models.Driver, error) {
	rows, err := q.db.QueryContext(ctx, listDriversQuery, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []models.Driver{}
	for rows.Next() {
		var i models.Driver
		if err := rows.Scan(
			&i.ID,
			&i.Username,
			&i.HashedPassword,
			&i.FullName,
			&i.Gender,
			&i.Balance,
			&i.Email,
			&i.CurrentCabID,
			&i.JoinedAt,
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

// updateDriverBalance is query for update user balance with id
const updateDriverBalance = `
UPDATE drivers 
SET balance = $2
WHERE id = $1
RETURNING id, username, hashed_password, full_name, gender, balance, email, current_cab_id, joined_at
`

// UpdateDriverBalance updates driver balance with UpdateDriverBalanceParams
func (q *PostgresStore) UpdateDriverBalance(ctx context.Context, arg models.UpdateDriverBalanceParams) (models.Driver, error) {
	row := q.db.QueryRowContext(ctx, updateDriverBalance, arg.ID, arg.Balance)
	var i models.Driver
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.HashedPassword,
		&i.FullName,
		&i.Gender,
		&i.Balance,
		&i.Email,
		&i.CurrentCabID,
		&i.JoinedAt,
	)
	return i, err
}

// updateDriverCurrentCabQuery is query for update driver current cab with UpdateDriverCurrentCabParams
const updateDriverCurrentCabQuery = `
UPDATE drivers 
SET current_cab_id = $2
WHERE id = $1
RETURNING id, username, hashed_password, full_name, gender, balance, email, current_cab_id, joined_at
`

func (q *PostgresStore) UpdateDriverCurrentCab(ctx context.Context, arg models.UpdateDriverCurrentCabParams) (models.Driver, error) {
	row := q.db.QueryRowContext(ctx, updateDriverCurrentCabQuery, arg.ID, arg.CurrentCabID)
	var i models.Driver
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.HashedPassword,
		&i.FullName,
		&i.Gender,
		&i.Balance,
		&i.Email,
		&i.CurrentCabID,
		&i.JoinedAt,
	)
	return i, err
}

// addDriverBalanceQuery is query for add to driver balacne
const addDriverBalanceQuery = `
UPDATE drivers 
SET balance = balance + $1
WHERE id = $2
RETURNING id, username, hashed_password, full_name, gender, balance, email, current_cab_id, joined_at
`

// AddDriverBalance Increases or decreases drier balance
func (q *PostgresStore) AddDriverBalance(ctx context.Context, arg models.AddDriverBalanceParams) (models.Driver, error) {
	row := q.db.QueryRowContext(ctx, addDriverBalanceQuery, arg.Amount, arg.ID)
	var i models.Driver
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.HashedPassword,
		&i.FullName,
		&i.Gender,
		&i.Balance,
		&i.Email,
		&i.CurrentCabID,
		&i.JoinedAt,
	)
	return i, err
}

// deleteDriverQuery is query for delete driver from database by id
const deleteDriverQuery = `
DELETE FROM drivers WHERE id = $1
`

// DeleteDriver deletes user from database by id
func (q *PostgresStore) DeleteDriver(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteDriverQuery, id)
	return err
}
