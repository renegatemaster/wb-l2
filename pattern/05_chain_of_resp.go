package pattern

import "fmt"

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/

/*
	Цепочка вызовов — поведенческий паттерн, позволяющий передавать выполнение запросов по цепочке
	Обработчик вызывается после вызова процесса предыдущим обработчиком

	Плюсы:
	- Каждый обработчик независимо выполняет свою роль
	- Принцип единственной обязанности
	- Принцип open/closed

	Минусы:
	- Запрос может остаться не обработанным
*/

type Service interface {
	Execute(*Data)
	SetNext(Service)
}

type Data struct {
	GetSource    bool
	UpdateSource bool
}

// Устройство, которое передаёт источник данных
type Device struct {
	Name string
	Next Service
}

func (device *Device) Execute(data *Data) {
	if data.GetSource {
		fmt.Printf("Данные были получены ранее от [%s]\n", device.Name)
		device.Next.Execute(data)
		return
	}
	fmt.Printf("Данные успешно получены [%s]\n", device.Name)
	data.GetSource = true
	device.Next.Execute(data)
}

func (device *Device) SetNext(svc Service) {
	device.Next = svc
}

// Сервис обновления данных, переданых устройством
type UpdateDataService struct {
	Name string
	Next Service
}

func (upd *UpdateDataService) Execute(data *Data) {
	if data.UpdateSource {
		fmt.Printf("Данные уже были обновлены ранее с помощью [%s]\n", upd.Name)
		upd.Next.Execute(data)
		return
	}
	fmt.Printf("Данные успешно обновлены [%s]\n", upd.Name)
	data.UpdateSource = true
	upd.Next.Execute(data)
}

func (upd *UpdateDataService) SetNext(svc Service) {
	upd.Next = svc
}

// Сервис сохранения данных
type DataService struct {
	Next Service
}

func (ds *DataService) Execute(data *Data) {
	if !data.UpdateSource {
		fmt.Println("Данные ещё не обработаны")
		return
	}
	fmt.Println("Данные сохранены")
}

func (ds *DataService) SetNext(svc Service) {
	ds.Next = svc
}

// Ниже код для проверки работы паттерна

// func main() {
// 	device := &Device{Name: "Device-01"}
// 	updateSvc := &UpdateDataService{Name: "Update-01"}
// 	dataSvc := &DataService{}

// 	device.SetNext(updateSvc)
// 	updateSvc.SetNext(dataSvc)

// 	data := &Data{}
// 	device.Execute(data)
// }
