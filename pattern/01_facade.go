package pattern

import (
	"errors"
	"fmt"
)

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

/*
	Суть паттерна "фасад" — реализовать простой доступ к сложной системе,
	скрыть лишнее и избавить пользователя от необходимости глубоко разбираться
	в устройстве системы.

	Плюсы:
	- Сокрытие сложной логики и простой интерфейс

	Минусы:
	- Интерфейс станет супер-объектом, к которому всё будет привязано

	Реализуем паттерн на примере магазина
	Пример взят с Ютуб-канала Dev Drift
	https://www.youtube.com/watch?v=rgnvjitzMdU
*/

type User struct {
	Name string
	Card *Card
}

func (user *User) GetBalance() float64 {
	return user.Card.Balance
}

type Card struct {
	Name    string
	Balance float64
	Bank    *Bank
}

func (card *Card) CheckBalance() error {
	fmt.Printf("[%s] Запрос в банк для проверки остатка %s\n", card.Name, card.Bank.Name)
	return card.Bank.CheckBalance(card.Name)
}

type Bank struct {
	Name  string
	Cards []Card
}

func (bank *Bank) CheckBalance(cardName string) error {
	fmt.Printf("[%s] Получение остатка по карте [%s]\n", bank.Name, cardName)
	for _, card := range bank.Cards {
		if card.Name != cardName {
			continue
		}
		if card.Balance <= 0 {
			return errors.New("на карте недостаточно средств")
		}
		fmt.Println("Баланс положительный")
		return nil
	}
	return errors.New("нет такой карты")
}

type Product struct {
	Name  string
	Price float64
}

type Shop struct {
	Name     string
	Products []Product
}

/*
В нашем примере этот метод и есть фасад
Пользователю нужно лишь назвать себя и товар,
Вся логика происходит внутри метода
*/
func (shop *Shop) Sell(user User, product string) error {

	fmt.Printf("[%s] Проверка баланса пользователя [%s]\n", shop.Name, user.Name)
	err := user.Card.CheckBalance()
	if err != nil {
		return err
	}

	fmt.Printf("[%s] Проверка может ли [%s] купить [%s]\n", shop.Name, user.Name, product)
	for _, prod := range shop.Products {
		if prod.Name != product {
			continue
		}
		if prod.Price > user.GetBalance() {
			return errors.New("на карте недостаточно средств для покупки")
		}
		fmt.Printf("[%s] был куплен [%s]\n", prod.Name, user.Name)
		return nil
	}
	return errors.New("товар не найден")
}

// Ниже код для проверки работы кода

// func main() {

// 	var (
// 		bank = Bank{
// 			Name:  "БАНК",
// 			Cards: []Card{},
// 		}
// 		card1 = Card{
// 			Name:    "Card-01",
// 			Balance: 200,
// 			Bank:    &bank,
// 		}
// 		card2 = Card{
// 			Name:    "Card-02",
// 			Balance: 5,
// 			Bank:    &bank,
// 		}
// 		user1 = User{
// 			Name: "Пользователь-01",
// 			Card: &card1,
// 		}
// 		user2 = User{
// 			Name: "Пользователь-02",
// 			Card: &card2,
// 		}
// 		prod = Product{
// 			Name:  "Хлеб",
// 			Price: 100,
// 		}
// 		shop = Shop{
// 			Name: "Seven/eleven",
// 			Products: []Product{
// 				prod,
// 			},
// 		}
// 	)

// 	fmt.Printf("[%s] Выпускает карт\n", bank.Name)
// 	bank.Cards = append(bank.Cards, card1, card2)

// 	err := shop.Sell(user1, "Хлеб")
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	err = shop.Sell(user2, "Хлеб")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// }
