package pkg

import "fmt"

// Реализация комьютера

type PersonalComputer struct {
	Type    string
	Core    int
	Memory  int
	Monitor bool
}

func NewPersonalComputer() Computer {
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
	fmt.Printf("%s: Ядра:[%d], Оперативная память: [%d], Наличие монитора: [%v]\n", pc.Type, pc.Core, pc.Memory, pc.Monitor)
}
