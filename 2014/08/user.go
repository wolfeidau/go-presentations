package main

type User struct {
	Name, Email string
}

func NewUser() *User { return &User{} }
