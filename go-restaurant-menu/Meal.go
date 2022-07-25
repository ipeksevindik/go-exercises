package main

type Meal struct {
	Name  string
	Price int
}



func NewMeal(name string, price int) *Meal {
	return &Meal{
		Name:  name,
		Price: price,
	}
}


