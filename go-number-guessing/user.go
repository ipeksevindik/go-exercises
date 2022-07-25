package main

import (
	"fmt"
	"strconv"
)

type User struct {
	Name  string
	Score int
}

func NewUser(name string) *User {
	return &User{
		Name:  name,
		Score: 100,
	}
}

func (user *User) Guess(realNum int) bool {
	fmt.Println("[", user.Name, "]", "Tahmininizi girin:")

	var guess string
	fmt.Scan(&guess)

	num, err := strconv.Atoi(guess)
	if err != nil {
		fmt.Println("Bir sayı girmediniz !")
		return false
	}

	user.Score -= 10
	if realNum < num {
		fmt.Println("Daha az bir sayı girin.")
		return false
	}
	if realNum > num {
		fmt.Println("Daha yüksek bir sayı girin.")
		return false
	}
	user.WinMessage()
	return true
}

func (user *User) WinMessage() {
	fmt.Println("----------------------------------------------------------")
	fmt.Println("Bravo ", user.Name, "sayıyı başarıyla tahmin ettin !!!")
	fmt.Println("Puanınız:", user.Score)
	fmt.Println("----------------------------------------------------------")
}
