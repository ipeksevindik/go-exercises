package main

import (
	"fmt"
	"math/rand"
	"time"
)



func main() {
	var userCount int
	fmt.Println("Oynayacak oyuncu sayısını girin: ")
	fmt.Scan(&userCount)

	realNum := randInt(0, 100*userCount)

	fmt.Println("0 ile ", 100*userCount, " arasında sayı bulma oyunu !!!!")

	users := []*User{}

	var isim string
	for i := 1; i <= userCount; i++ {
		fmt.Println(i, ". Oyuncu isminizi girin: ")
		fmt.Scan(&isim)
		users = append(users, NewUser(isim))
	}

	idx := 0
	for {
		currentUser := users[idx]

		if currentUser.Guess(realNum) {
			break
		}
		if currentUser.Score <= 0 {
			fmt.Println("[", currentUser.Name, "]", " kaybettiniz :(")
			users = append(users[:idx], users[idx+1:]...)
		}
		idx += 1
		if idx >= len(users) {
			idx = 0
		}
	}

}

func randInt(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min)
}
