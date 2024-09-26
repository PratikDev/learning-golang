package main

import "fmt"

type User struct {
	Name       string
	Email      string
	Age        int
	IsVerified bool
}

func main() {
	user := User{Name: "pratik dev", Email: "pratikdevofficial1@gmail.com", Age: 22, IsVerified: true}
	fmt.Printf("%+v\n", user)
}
