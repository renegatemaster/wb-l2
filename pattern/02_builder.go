package pattern

import "fmt"

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

/*
	"Строитель" — порождающий паттерн проектирования.
	Позволяет создавать сложные объекты, используя "шаги" —
	— маленькие промежуточные объекты со своей простой логикой

	Плюсы:
	- Позволяет пошагово создать сложный продукт
	- Позволяет использовать один и тот же код для создания различных объектов
	- Изолирует сложную логику объекта

	Минусы:
	- Усложняет код из-за введения дополнительных объектов
	- Привязка к конкретному объекту строителя

	Реализуем паттерн на примере сборки компьютера
	Пример взят с канала Dev Drift
	https://www.youtube.com/watch?v=_4dlAjE2rrc&list=PLxj7Nz8YYkVUlLTvfkS1Q8OW8_rM11PoS&index=38
*/

const (
	AsusCollectorType = "asus"
	HpCollectorType   = "hp"
)

type Collector interface {
	SetCore()
	SetBrand()
	SetMemory()
	SetMonitor()
	SetGraphicCard()
	GetComputer() Computer
}

func GetCollector(collectorType string) Collector {
	switch collectorType {
	default:
		return nil
	case AsusCollectorType:
		return &AsusCollector{}
	case HpCollectorType:
		return &HpCollector{}
	}
}

// ASUS

type AsusCollector struct {
	Core        int
	Brand       string
	Memory      int
	Monitor     int
	GraphicCard int
}

func (collector *AsusCollector) SetCore() {
	collector.Core = 4
}

func (collector *AsusCollector) SetBrand() {
	collector.Brand = "Asus"
}

func (collector *AsusCollector) SetMemory() {
	collector.Memory = 8
}

func (collector *AsusCollector) SetMonitor() {
	collector.Monitor = 1
}

func (collector *AsusCollector) SetGraphicCard() {
	collector.GraphicCard = 1
}

func (collector *AsusCollector) GetComputer() Computer {
	return Computer{
		Core:        collector.Core,
		Brand:       collector.Brand,
		Memory:      collector.Memory,
		Monitor:     collector.Monitor,
		GraphicCard: collector.GraphicCard,
	}
}

// HP

type HpCollector struct {
	Core        int
	Brand       string
	Memory      int
	Monitor     int
	GraphicCard int
}

func (collector *HpCollector) SetCore() {
	collector.Core = 4
}

func (collector *HpCollector) SetBrand() {
	collector.Brand = "Hp"
}

func (collector *HpCollector) SetMemory() {
	collector.Memory = 16
}

func (collector *HpCollector) SetMonitor() {
	collector.Monitor = 2
}

func (collector *HpCollector) SetGraphicCard() {
	collector.GraphicCard = 1
}

func (collector *HpCollector) GetComputer() Computer {
	return Computer{
		Core:        collector.Core,
		Brand:       collector.Brand,
		Memory:      collector.Memory,
		Monitor:     collector.Monitor,
		GraphicCard: collector.GraphicCard,
	}
}

type Computer struct {
	Core        int
	Brand       string
	Memory      int
	Monitor     int
	GraphicCard int
}

func (pc *Computer) Print() {
	fmt.Printf(
		"PC '%s': Core[%d], Mem[%d], Monitor[%d], GraphicCard[%d]\n",
		pc.Brand, pc.Core, pc.Memory, pc.Monitor, pc.GraphicCard,
	)
}

type Factory struct {
	Collector Collector
}

func NewFactory(collector Collector) *Factory {
	return &Factory{Collector: collector}
}

func (factory *Factory) SetCollector(collector Collector) {
	factory.Collector = collector
}

func (factory *Factory) CreateComputer() Computer {
	factory.Collector.SetCore()
	factory.Collector.SetBrand()
	factory.Collector.SetMemory()
	factory.Collector.SetMonitor()
	factory.Collector.SetGraphicCard()
	return factory.Collector.GetComputer()
}

// Код ниже для проверки работы паттерна

// func main() {
// 	asusCollector := GetCollector("asus")
// 	hpCollector := GetCollector("hp")

// 	// Производим на фабрике одни компьютеры
// 	factory := NewFactory(asusCollector)
// 	asusComputer := factory.CreateComputer()
// 	asusComputer.Print()

// 	// А теперь другие на той же фабрике
// 	factory.SetCollector(hpCollector)
// 	hpComputer := factory.CreateComputer()
// 	hpComputer.Print()
// }
