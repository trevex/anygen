package main

import (
	"context"
	"fmt"
)

type HelloResponse struct {
	Name string
}

//anygen:api public method=GET path=/hello/:name
func Hello(ctx context.Context, name string) (*HelloResponse, error) {
	return &HelloResponse{name}, nil
}

func main() {
	fmt.Println("TODO")
}
