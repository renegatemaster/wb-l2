package pattern

import "fmt"

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/

/*
	"Команда" — поведенческий паттерн, который инкапсулирует запрос в виде объекта
	Клиент через инвокер взаимодействует с получателем
	Инвокер вызывает по нажатию кнопки определённую команду, которая знает своего получателя

	Плюсы:
	- Добавление новых команд без изменения кода
	- Упрощает тестирование кода

	Минусы:
	- Усложнение кода за счёт увеличения количества объектов
	- В некоторых случаях отмена команды может быть нерелевантной, что требует доп. обработки
*/

type Command interface {
	Execute()
	Undo()
}

// Receiver
type TV struct{}

func (tv *TV) On() {
	fmt.Println("TV is on")
}

func (tv *TV) Off() {
	fmt.Println("TV is off")
}

// Concrete command
type TVCommand struct {
	tv TV
}

func (tvc TVCommand) Execute() {
	tvc.tv.On()
}

func (tvc TVCommand) Undo() {
	tvc.tv.Off()
}

// Another receiver
type Microwave struct{}

func (mw *Microwave) StartHeating() {
	fmt.Println("Microwave is heating food")
}

func (mw *Microwave) StopHeating() {
	fmt.Println("Microwave stopped heating")
}

// Another concrete command
type MWCommand struct {
	mw Microwave
}

func (mwc MWCommand) Execute() {
	mwc.mw.StartHeating()
}

func (mwc MWCommand) Undo() {
	mwc.mw.StopHeating()
}

// Invoker
type Pult struct {
	Command Command
}

func (pult *Pult) SetCommand(command Command) {
	pult.Command = command
}

func (pult *Pult) PressButton() {
	pult.Command.Execute()
}

func (pult *Pult) PressUndo() {
	pult.Command.Undo()
}

// Ниже код для тестирования паттерна

// func main() {
// 	var (
// 		tv     TV
// 		mw     Microwave
// 		commTV TVCommand
// 		commMW MWCommand
// 		pult   Pult
// 	)

// 	commTV.tv = tv
// 	commMW.mw = mw

// 	pult.SetCommand(commTV)

// 	pult.PressButton()
// 	pult.PressUndo()

// 	pult.SetCommand(commMW)

// 	pult.PressButton()
// 	pult.PressUndo()
// }
