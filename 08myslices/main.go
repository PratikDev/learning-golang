package main

import "fmt"

func main() {
	users := []string{"user1", "user2", "user3"}

	users = append(users, "user4", "user5")
	fmt.Println("Users are:", users)

	users = append(users[3:6]) // [3:6] => from 3rd index till 6th (exclusive)
	fmt.Println("Users are:", users)

	users = append(users[1:]) // [1:] => from 1st index till the end (exclusive)
	fmt.Println("Users are:", users)
}
