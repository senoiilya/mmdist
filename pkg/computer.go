package pkg

import "fmt"

// константы наших типов

const (
	ServerType           = "server"
	PersonalComputerType = "personal"
	NotebookType         = "notebook"
)

type Computer interface {
	GetType() string // тип
	PrintDetails()   // вывод детализированной информации о самом объекте
	String() string
}

// Фабричный метод, который будет инициализировать структуры

func New(typeName string) Computer {
	switch typeName {
	default:
		fmt.Printf("%s: Несуществующий тип объекта!\n", typeName)
		return nil
	case ServerType:
		return NewServer()
	case PersonalComputerType:
		return NewPersonalComputer()
	case NotebookType:
		return NewNotebook()
	}

}
