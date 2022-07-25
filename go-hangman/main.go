package main

import "fmt"

func main() {

	games := []*Game{
		NewOyun("mubasir"),
		NewOyun("ahiret"),
		NewOyun("makas"),
		NewOyun("su"),
		NewOyun("elveda"),
		NewOyun("klavye"),
		NewOyun("gozluk"),
		NewOyun("samanyolu"),
		NewOyun("atlet"),
		NewOyun("pantolon"),
		NewOyun("masa"),
		NewOyun("samet"),
		NewOyun("sandalye"),
		NewOyun("ipek"),
		NewOyun("merve"),
		NewOyun("bugra"),
		NewOyun("adil"),
		NewOyun("cagri"),
		NewOyun("mugla"),
		NewOyun("istanbul"),
	}

	game :=  ChooseGame(games)

	var ch string

	for {
		game.KelimeyiYazdir()
		fmt.Scan(&ch)
		game.Guess(ch)
		if game.End() {
			break
		}
	}
}
