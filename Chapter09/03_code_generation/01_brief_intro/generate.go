package _1_brief_intro

import "context"

//go:generate mockery --name=UserLoader --case=underscore --testonly --inpackage
type UserLoader interface {
	LoadByID(ctx context.Context, userID string) (*User, error)
}

type User struct {
	ID   string
	Name string
}
