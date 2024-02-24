package pattern

import "fmt"

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/

/*
	Фабричный метод — порождающий паттерн проектирования
	Определяет общий интерфейс поведения для создаваемых объектов

	Плюсы:
	- Избавляет от привязки к конкретному типу объекта
	- Общий конструктор создания, что упрощает добавлений новых объектов от базового
	- Принцип open/closed

	Минусы:
	- Может привести к созданию больших параллельных иерархий объектов
	- Появляется "божественный" конструктор, к нему происходит привязка
*/

const (
	ServerType           = "server"
	PersonalComputerType = "personal"
	NotebookType         = "notebook"
)

type Machine interface {
	GetType() string
	PrintDetails()
}

// Фабричный метод
func New(typeName string) Machine {
	switch typeName {
	default:
		fmt.Printf("Не найден тип объекта [%s]\n", typeName)
		return nil
	case ServerType:
		return NewServer()
	case PersonalComputerType:
		return NewPersonalComputer()
	case NotebookType:
		return NewNotebook()
	}

}

// Server
type Server struct {
	Type   string
	Core   int
	Memory int
}

func NewServer() Machine {
	return Server{
		Type:   ServerType,
		Core:   16,
		Memory: 256,
	}
}

func (pc Server) GetType() string {
	return pc.Type
}

func (pc Server) PrintDetails() {
	fmt.Printf("%s Core[%d], Mem[%d]\n", pc.Type, pc.Core, pc.Memory)
}

// PC
type PersonalComputer struct {
	Type    string
	Core    int
	Memory  int
	Monitor bool
}

func NewPersonalComputer() Machine {
	return PersonalComputer{
		Type:    PersonalComputerType,
		Core:    8,
		Memory:  16,
		Monitor: true,
	}
}

func (pc PersonalComputer) GetType() string {
	return pc.Type
}

func (pc PersonalComputer) PrintDetails() {
	fmt.Printf("%s Core[%d], Mem[%d], Monitor[%v]\n", pc.Type, pc.Core, pc.Memory, pc.Monitor)
}

// Notebook
type Notebook struct {
	Type    string
	Core    int
	Memory  int
	Monitor bool
}

func NewNotebook() Machine {
	return Notebook{
		Type:    NotebookType,
		Core:    4,
		Memory:  8,
		Monitor: true,
	}
}

func (pc Notebook) GetType() string {
	return pc.Type
}

func (pc Notebook) PrintDetails() {
	fmt.Printf("%s Core[%d], Mem[%d], Monitor[%v]\n", pc.Type, pc.Core, pc.Memory, pc.Monitor)
}

// Ниже код для проверки работы паттерна

// func main() {

// 	types := []string{ServerType, PersonalComputerType, NotebookType, "Phone"}

// 	for _, typeName := range types {
// 		computer := New(typeName)
// 		if computer == nil {
// 			continue
// 		}
// 		computer.PrintDetails()
// 	}
// }
