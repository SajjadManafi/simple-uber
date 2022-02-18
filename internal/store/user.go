package store

import (
	"context"

	"github.com/SajjadManafi/simple-uber/models"
)

// createUserQuery is query for insert user
const createUserQuery = `
INSERT INTO users (
  username,
  hashed_password,
  full_name,
  gender,
  balance,
  email
) VALUES (
  $1, $2, $3, $4, 0, $5
) RETURNING id, username, hashed_password, full_name, gender, balance, email, joined_at
`

// CreateUser creates user in database
func (q *PostgresStore) CreateUser(ctx context.Context, arg models.CreateUserParams) (models.User, error) {
	row := q.db.QueryRowContext(ctx, createUserQuery,
		arg.Username,
		arg.HashedPassword,
		arg.FullName,
		arg.Gender,
		arg.Email,
	)
	var i models.User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.HashedPassword,
		&i.FullName,
		&i.Gender,
		&i.Balance,
		&i.Email,
		&i.JoinedAt,
	)
	return i, err
}

// getUserQuery is query for get user
const getUserQuery = `
SELECT id, username, hashed_password, full_name, gender, balance, email, joined_at FROM users
WHERE id = $1 LIMIT 1
`

// GetUser gets user from database
func (q *PostgresStore) GetUser(ctx context.Context, id int32) (models.User, error) {
	row := q.db.QueryRowContext(ctx, getUserQuery, id)
	var i models.User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.HashedPassword,
		&i.FullName,
		&i.Gender,
		&i.Balance,
		&i.Email,
		&i.JoinedAt,
	)
	return i, err
}

// listUsersQuery is query for list users
const listUsersQuery = `
SELECT id, username, hashed_password, full_name, gender, balance, email, joined_at FROM users
ORDER BY id
LIMIT $1
OFFSET $2
`

// ListUsers gets users and return slice of users (with limit and offset)
func (q *PostgresStore) ListUsers(ctx context.Context, arg models.ListUsersParams) ([]models.User, error) {
	rows, err := q.db.QueryContext(ctx, listUsersQuery, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []models.User{}
	for rows.Next() {
		var i models.User
		if err := rows.Scan(
			&i.ID,
			&i.Username,
			&i.HashedPassword,
			&i.FullName,
			&i.Gender,
			&i.Balance,
			&i.Email,
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

// updateUserQuery is query for update user
const updateUserQuery = `
UPDATE users 
SET balance = $2
WHERE id = $1
RETURNING id, username, hashed_password, full_name, gender, balance, email, joined_at
`

// UpdateUser updated user balance with user id
func (q *PostgresStore) UpdateUser(ctx context.Context, arg models.UpdateUserParams) (models.User, error) {
	row := q.db.QueryRowContext(ctx, updateUserQuery, arg.ID, arg.Balance)
	var i models.User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.HashedPassword,
		&i.FullName,
		&i.Gender,
		&i.Balance,
		&i.Email,
		&i.JoinedAt,
	)
	return i, err
}

// deleteUserQuery is query for delete user
const deleteUserQuery = `
DELETE FROM users WHERE id = $1
`

// DeleteUser deletes user from database
func (q *PostgresStore) DeleteUser(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteUserQuery, id)
	return err
}

// addUserBalanceQuery is query for add to user balacne
const addUserBalanceQuery = `-- name: AddUserBalance :one
UPDATE users 
SET balance = balance + $1
WHERE id = $2
RETURNING id, username, hashed_password, full_name, gender, balance, email, joined_at
`

// AddUserBalance Increases or decreases user balance
func (q *PostgresStore) AddUserBalance(ctx context.Context, arg models.AddUserBalanceParams) (models.User, error) {
	row := q.db.QueryRowContext(ctx, addUserBalanceQuery, arg.Amount, arg.ID)
	var i models.User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.HashedPassword,
		&i.FullName,
		&i.Gender,
		&i.Balance,
		&i.Email,
		&i.JoinedAt,
	)
	return i, err
}
