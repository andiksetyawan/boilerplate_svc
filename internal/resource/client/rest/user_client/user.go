package user_client

import "context"

type User struct {
	ID   int
	Name string
}

type UserClient interface {
	GetByID(ctx context.Context, userID int) (User, error)
}

//TODO add implementation
