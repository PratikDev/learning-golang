package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	fmt.Println(currentTime.Format(time.RFC1123Z)) // ref: https://pkg.go.dev/time#pkg-constants

	createdTime := time.Date(2020, time.January, 14, 13, 14, 15, 155, time.Local)
	fmt.Println(createdTime.Format(time.RFC1123Z))
}
