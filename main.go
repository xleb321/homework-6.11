package main

import (
	"fmt"
	"sync"
)

type Products interface {
	Add()
	Delete()
	PrintAll()
	PrintByCategory()
}

// Под него, реализуй всё из Products
type Product struct {
	name     string
	category string
}

// что это за тип, ... буду у тебя спрашивать на уроке
// используй его (1)
// type ProductList []Product

// Что-то вроде такого (2)
// func (p Products) PrintByCategory(*wg, cat string) {
// 	defer wg.Done()
// 	fmt.Printf("%s : ", cat)
// 	for _, i := range allProducts {
// 		if i.category == cat {
// 			fmt.Printf("%s ", i.name)
// 		}
// 	}
// }

func main() {
	var wg sync.WaitGroup // (3)
	var allProducts = []Product{{"Банан", "Фрукты"}, {"Помидор", "Овощи"}, {"Яблоко", "Фрукты"}, {"Груша", "Фрукты"}, {"Морковь", "Овощи"}}
	//                 ^ тут (1)

	wg.Add(2)
	
	go allProducts.PrintByCategory("Фрукты", &wg)
	// Для "Овощи" соотвецтвенно

	wg.Wait()
}

// review
