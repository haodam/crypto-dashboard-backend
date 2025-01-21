// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package postgres

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type User struct {
	UserID      int64              `json:"user_id"`
	Username    string             `json:"username"`
	FullName    string             `json:"full_name"`
	Gender      pgtype.Text        `json:"gender"`
	Email       string             `json:"email"`
	PhoneNumber pgtype.Text        `json:"phone_number"`
	DateOfBirth pgtype.Date        `json:"date_of_birth"`
	Address     pgtype.Text        `json:"address"`
	CreatedAt   time.Time          `json:"created_at"`
	UpdatedAt   pgtype.Timestamptz `json:"updated_at"`
	IsActive    pgtype.Bool        `json:"is_active"`
	Role        pgtype.Text        `json:"role"`
}