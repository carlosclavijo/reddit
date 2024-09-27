package repository

import "github.com/carlosclavijo/reddit/internal/models"

type DatabaseRepo interface {
	AllUsers() bool
	InsertUser(res models.User) error
}
