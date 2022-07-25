package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func PrintMenu(menu []*Meal) {
	fmt.Println("--------------------------------------------------")
	fmt.Println("Seçim yapınız:")

	for i, meal := range menu {
		fmt.Println(i+1, "-", meal.Name, "Price:", meal.Price)
	}
	fmt.Println("--------------------------------------------------")
}

func main() {

	menu := []*Meal{
		NewMeal("French Fries", 5),
		NewMeal("Hamburger", 20),
		NewMeal("Pizza", 30),
	}

	var order int
	wallet := 0

	fmt.Println("Cüzdanınıza ekleyeceğiniz Tutarı giriniz:")
	fmt.Scan(&wallet)

	for {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
		PrintMenu(menu)

		fmt.Scan(&order)
		if order == 0 {
			break
		}

		order -= 1

		wallet -= menu[order].Price

		if wallet > 0 {
			fmt.Println(menu[order].Name, " sepete eklediniz. Fiyatı: ", menu[order].Price)
		} else {
			wallet = 0
			fmt.Println("Paranız bitti :(")
			break
		}

		time.Sleep(time.Second * 1)
	}
	fmt.Println("Cuzdanınızda kalan miktar:", wallet)

}
