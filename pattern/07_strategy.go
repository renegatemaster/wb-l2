package pattern

import "fmt"

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/

/*
	Стратегия — поведенческий паттерн проектирования
	Определяет схожие алгоритмы и помещает каждый в свою отдельную структуру

	Плюсы:
	- Замена алгоритмов налету
	- Изолирует код и данные алгоритмов от остальных объектов
	- Уход от наследования
	- Принцип open/closed

	Минусы:
	- Усложнение программы большим количеством кода
	- Клиент должен знать разницу между стратегиями для выбора подходящей
*/

type Strategy interface {
	Route(startPoint int, endPoint int)
}

type Navigator struct {
	Strategy
}

func (nav *Navigator) SetStrategy(str Strategy) {
	nav.Strategy = str
}

// Road
type RoadStrategy struct {
}

func (r *RoadStrategy) Route(startPoint, endPoint int) {
	avgSpeed := 30
	avgTime := 40
	trafficJam := 2
	totalDistance := endPoint - startPoint
	totalTime := totalDistance * avgTime * trafficJam

	fmt.Printf(
		"Road: \nA[%d] to B[%d]: Avg speed [%d], Traffic jam [%d], Total distance [%d], Total time [%d]\n",
		startPoint, endPoint, avgSpeed, trafficJam, totalDistance, totalTime,
	)
}

// Public transport
type PublicTransportStrategy struct {
}

func (r *PublicTransportStrategy) Route(startPoint, endPoint int) {
	avgSpeed := 40
	avgTime := 40
	totalDistance := endPoint - startPoint
	totalTime := totalDistance * avgTime

	fmt.Printf(
		"PublicTransport: \nA[%d] to B[%d]: Avg speed [%d], Total distance [%d], Total time [%d]\n",
		startPoint, endPoint, avgSpeed, totalDistance, totalTime,
	)
}

// Walk
type WalkStrategy struct {
}

func (r *WalkStrategy) Route(startPoint, endPoint int) {
	avgSpeed := 4
	avgTime := 60
	totalDistance := endPoint - startPoint
	totalTime := totalDistance * avgTime

	fmt.Printf(
		"Walk: \nA[%d] to B[%d]: Avg speed [%d], Total distance [%d], Total time [%d]\n",
		startPoint, endPoint, avgSpeed, totalDistance, totalTime,
	)
}

// Ниже код для проверки работы паттерна

// func main() {

// 	var (
// 		start      = 10
// 		end        = 100
// 		strategies = []Strategy{
// 			&RoadStrategy{},
// 			&PublicTransportStrategy{},
// 			&WalkStrategy{},
// 		}
// 	)

// 	nav := Navigator{}
// 	for _, strategy := range strategies {
// 		nav.SetStrategy(strategy)
// 		nav.Route(start, end)
// 	}
// }
