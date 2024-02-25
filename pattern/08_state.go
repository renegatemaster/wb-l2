package pattern

import (
	"errors"
	"fmt"
)

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/

/*
	Состояние — поведенческий паттерн проектирования
	Позволяет объектам менять поведение в зависимости от своего состояния
	Извне кажется, что объект изменился полностью
	Программа может находиться в одном из нескольких состояний, которые постоянно сменяют друг друга
	Количество этих состояний и переходов между ними конечно
	Находясь в разных состояниях программа может по-разному реагировать на одни и те же события

	Плюсы:
	- Избавляет от множества условных операторов
	- Концентрирует в одном месте логику определённого состояния

	Минусы:
	- Может неоправданно усложнить код, если состояний мало и они редко меняются
*/

type State interface {
	AddItem(int) error
	RequestItem() error
	InsertMoney(money int) error
	DispenseItem() error
}

type VendingMachine struct {
	HasItem       State
	ItemRequested State
	HasMoney      State
	NoItem        State
	CurrentState  State
	ItemCount     int
	ItemPrice     int
}

func NewVendingMachine(itemCount, itemPrice int) *VendingMachine {
	v := &VendingMachine{
		ItemCount: itemCount,
		ItemPrice: itemPrice,
	}

	hasItemState := &HasItemState{
		VendingMachine: v,
	}
	itemRequestedState := &ItemRequestedState{
		VendingMachine: v,
	}
	hasMoneyState := &HasMoneyState{
		VendingMachine: v,
	}
	noItemState := &NoItemState{
		VendingMachine: v,
	}
	v.HasItem = hasItemState
	v.ItemRequested = itemRequestedState
	v.HasMoney = hasMoneyState
	v.NoItem = noItemState
	v.SetState(hasItemState)

	return v
}

func (v *VendingMachine) AddItem(count int) error {
	return v.CurrentState.AddItem(count)
}

func (v *VendingMachine) RequestItem() error {
	return v.CurrentState.RequestItem()
}

func (v *VendingMachine) InsertMoney(money int) error {
	return v.CurrentState.InsertMoney(money)
}

func (v *VendingMachine) DispenseItem() error {
	return v.CurrentState.DispenseItem()
}

func (v *VendingMachine) SetState(s State) {
	v.CurrentState = s
}

func (v *VendingMachine) IncrementItemCount(count int) {
	v.ItemCount += count
}

// STATES

// No item
type NoItemState struct {
	VendingMachine *VendingMachine
}

func (s *NoItemState) AddItem(count int) error {
	s.VendingMachine.IncrementItemCount(count)
	fmt.Printf("%d items added \n", count)
	s.VendingMachine.SetState(s.VendingMachine.HasItem)
	return nil
}

func (s *NoItemState) RequestItem() error {
	return errors.New("item out of stock")
}

func (s *NoItemState) InsertMoney(money int) error {
	return errors.New("item out of stock")
}

func (s *NoItemState) DispenseItem() error {
	return errors.New("item out of stock")
}

// Requested
type ItemRequestedState struct {
	VendingMachine *VendingMachine
}

func (s *ItemRequestedState) AddItem(count int) error {
	return errors.New("item dispense in progress")
}

func (s *ItemRequestedState) RequestItem() error {
	return errors.New("item already requested")
}

func (s *ItemRequestedState) InsertMoney(money int) error {
	if money < s.VendingMachine.ItemPrice {
		return fmt.Errorf("received [%d], needed [%d]", money, s.VendingMachine.ItemPrice)
	}
	fmt.Println("Money received")
	s.VendingMachine.SetState(s.VendingMachine.HasMoney)
	return nil
}

func (s *ItemRequestedState) DispenseItem() error {
	return errors.New("please insert money first")
}

// Has money
type HasMoneyState struct {
	VendingMachine *VendingMachine
}

func (s *HasMoneyState) AddItem(count int) error {
	return errors.New("item dispense in progress")
}

func (s *HasMoneyState) RequestItem() error {
	return errors.New("item dispense in progress")
}

func (s *HasMoneyState) InsertMoney(money int) error {
	return errors.New("item dispense in progress")
}

func (s *HasMoneyState) DispenseItem() error {
	fmt.Println("Dispensing Item")
	s.VendingMachine.ItemCount -= 1
	if s.VendingMachine.ItemCount == 0 {
		s.VendingMachine.SetState(s.VendingMachine.NoItem)
	} else {
		s.VendingMachine.SetState(s.VendingMachine.HasItem)
	}
	return nil
}

// Has item
type HasItemState struct {
	VendingMachine *VendingMachine
}

func (s *HasItemState) AddItem(count int) error {
	fmt.Printf("%d items added\n", count)
	s.VendingMachine.IncrementItemCount(count)
	return nil
}

func (s *HasItemState) RequestItem() error {
	if s.VendingMachine.ItemCount == 0 {
		s.VendingMachine.SetState(s.VendingMachine.NoItem)
		return errors.New("no item present")
	}
	fmt.Println("Item requested")
	s.VendingMachine.SetState(s.VendingMachine.ItemRequested)
	return nil
}

func (s *HasItemState) InsertMoney(money int) error {
	return errors.New("please select item first")
}

func (s *HasItemState) DispenseItem() error {
	return errors.New("please select item first")
}

// Ниже код для проверки работы паттерна

// func main() {

// 	vendingMachine := NewVendingMachine(1, 10)

// 	err := vendingMachine.RequestItem()
// 	if err != nil {
// 		log.Fatalf(err.Error())
// 	}

// 	err = vendingMachine.InsertMoney(10)
// 	if err != nil {
// 		log.Fatalf(err.Error())
// 	}

// 	err = vendingMachine.DispenseItem()
// 	if err != nil {
// 		log.Fatalf(err.Error())
// 	}

// 	fmt.Println()

// 	err = vendingMachine.AddItem(2)
// 	if err != nil {
// 		log.Fatalf(err.Error())
// 	}

// 	fmt.Println()

// 	err = vendingMachine.RequestItem()
// 	if err != nil {
// 		log.Fatalf(err.Error())
// 	}

// 	err = vendingMachine.InsertMoney(10)
// 	if err != nil {
// 		log.Fatalf(err.Error())
// 	}

// 	err = vendingMachine.DispenseItem()
// 	if err != nil {
// 		log.Fatalf(err.Error())
// 	}
// }
