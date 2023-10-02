package main

import (
	"context"
	"fmt"
)

type User struct {
	ID string
}
type Pagination struct{}
type JSON[T any] struct {
	Value T
}
type Query[T any] struct {
	Value T
}

//anygen:api public method=GET path=/hello/:name
func Hello(ctx context.Context, name string) (*HelloResponse, error) {
	return &HelloResponse{name}, nil
}

//anygen:api public method=POST path=/user
func CreateUser(ctx context.Context, data JSON[User]) (*User, error) {
	return &data.Value, nil
}

//anygen:api public method=GET path=/users
func ListUsers(ctx context.Context, query Query[Pagination]) ([]User, error) {
	return nil, nil
}

//anygen:api public method=POST path=/user body=user
func CreateUser2(ctx context.Context, user User) (*User, error) {
	return &user, nil
}

//anygen:route ListUsers method=GET path=/users?offset=0&limit=10 header=X-Header=myheader
func listUsers(ctx context.Context, offset uint, limit uint) ([]User, error) {
	return nil, nil
}

//anygen:route ListUsers method=GET path=/users?offset=0&limit=10 header=X-Header=myheader

// TODO: FromRequestParts, IntoResponse

func main() {
	fmt.Println("TODO")
}
