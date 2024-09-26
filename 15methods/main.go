package main

import "fmt"

type User struct {
	Name       string
	Email      string
	Age        int
	IsVerified bool
}

func (u User) getVerificationStatus() {
	fmt.Println("User verfication status is:", u.IsVerified)
}

func (u *User) changeEmail(newMail string) {
	u.Email = newMail
	fmt.Println("New email is:", u.Email)
}

func main() {
	user := User{Name: "pratik dev", Email: "pratikdevofficial1@gmail.com", Age: 22, IsVerified: true}
	user.getVerificationStatus()
	user.changeEmail("test@pratik.dev")
	fmt.Println("Old email is:", user.Email)
}
