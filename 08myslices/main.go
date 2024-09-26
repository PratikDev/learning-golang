package main

import (
	"fmt"
)

func main() {
	users := []string{"user1", "user2", "user3"}

	users = append(users, "user4", "user5")
	fmt.Println("Users are:", users)

	users = append(users[3:6], "user6", "user7") // [3:6] => from 3rd index till 6th (exclusive)
	fmt.Println("Users are:", users)

	scores := make([]int, 2)
	scores[0] = 232
	scores[1] = 332
	scores = append(scores, 423, 129) // with `append`, it'll readjust the memory if needed
	fmt.Println("Scores are:", scores)

	courses := []string{"php", "js", "go", "py"}
	index := int(2)
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println("courses are:", courses)
}
