package main

import (
	crypto "crypto/rand"
	"fmt"
	"math/big"
	"math/rand"
	"time"
)

func main() {
	rand.New(rand.NewSource(time.Now().UnixMicro()))
	randomNumber := rand.Intn(5) + 1
	fmt.Println("Math random: ", randomNumber)

	cryptoRandomNumber, _ := crypto.Int(crypto.Reader, big.NewInt(5))
	result := new(big.Int).Add(cryptoRandomNumber, big.NewInt(1))
	fmt.Println("Crypto random: ", result)
}
