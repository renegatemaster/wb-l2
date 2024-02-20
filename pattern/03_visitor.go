package pattern

import "fmt"

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/

/*
	"Посетитель" добавляет новый функционал объекту без его изменения.

	Плюсы:
	- не вносим в структуру объекта множество не свяазанных между собой операции
	- не изменяем структуру объекта
	- спокойно добавляем новые операции
	- в одном месте делаем опрации над разными независимыми классами

	Минусы:
	- усложняет расширение иерархии классов, поскольку новые классы обычно требуют добавления нового метода visit для каждого посетителя.
*/

type Element interface {
	Accept(visitor Visitor)
}

type ConcreteElementA struct{}

func (ce *ConcreteElementA) Accept(visitor Visitor) {
	visitor.VisitConcreteElementA(ce)
}

type ConcreteElementB struct{}

func (ce *ConcreteElementB) Accept(visitor Visitor) {
	visitor.VisitConcreteElementB(ce)
}

type Visitor interface {
	VisitConcreteElementA(element *ConcreteElementA)
	VisitConcreteElementB(element *ConcreteElementB)
}

type ConcreteVisitor struct{}

func (cv *ConcreteVisitor) VisitConcreteElementA(element *ConcreteElementA) {
	fmt.Println("Visiting ConcreteElementA")
}

func (cv *ConcreteVisitor) VisitConcreteElementB(element *ConcreteElementB) {
	fmt.Println("Visiting ConcreteElementB")
}

// Код ниже для проверки работы паттерна

// func main() {
// 	visitor := &ConcreteVisitor{}

// 	elementA := &ConcreteElementA{}
// 	elementA.Accept(visitor)

// 	elementB := &ConcreteElementB{}
// 	elementB.Accept(visitor)
// }
