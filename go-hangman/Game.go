package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

type Game struct {
	Word            string
	Skor             int
	KnownCharacters map[int]bool
}

func NewOyun(word string) *Game {
	knownCharacters := map[int]bool{}

	for i := 0; i < len(word); i++ {
		knownCharacters[i] = false
	}

	return &Game{
		Word:          word,
		Skor:           len(word) * 5,
		KnownCharacters: knownCharacters,
	}
}

func (game *Game) End() bool {
	if game.Skor <= 0 {
		fmt.Println("OYUNU KAYBETTİNİZ !")
		return true
	}

	for _, known := range game.KnownCharacters {
		if !known {
			return false
		}
	}

	fmt.Println("TÜM HARFLERİ BULDUNUZ ! OYUN KAZANILDI")
	return true
}

func (game *Game) Guess(ch string) {
	hicBildikMi := false
	for i := 0; i < len(game.Word); i++ {
		harf := string(game.Word[i])
		if ch == harf {
			game.KnownCharacters[i] = true
			hicBildikMi = true
		}
	}
	if !hicBildikMi {
		game.Skor -= 5
	}
}

func (game *Game) KelimeyiYazdir() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
	fmt.Println("Şu anki canınız : ", game.Skor)
	for i := 0; i < len(game.Word); i++ {
		if game.KnownCharacters[i] {
			fmt.Print(string(game.Word[i]))
		} else {
			fmt.Print("_")
		}
	}
	fmt.Println()
}

func ChooseGame(words []*Game) *Game {
	return words[randInt(0, len(words))]

}
func randInt(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min)
}
