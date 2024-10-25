package domain

import "my_app/src/infrastructure"

// User

type User struct {
	Id    int
	Name  string
	Email string
}

type DUser = infrastructure.DataDecorator[User]

// other domain models
