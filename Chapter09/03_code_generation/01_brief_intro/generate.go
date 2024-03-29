package _1_brief_intro

import "context"

//go:generate mockery --name=UserLoader --case=underscore --testonly --inpackage
type UserLoader interface {
	LoadByID(ctx context.Context, userID int64) (*User, error)
}

type User struct {
	ID   int64
	Name string
}
