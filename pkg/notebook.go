package pkg

import "fmt"

// Реализация комьютера

type Notebook struct {
	Type    string
	Core    int
	Memory  int
	Monitor bool
}

func NewNotebook() Computer {
	return Notebook{
		Type:    NotebookType,
		Core:    4,
		Memory:  8,
		Monitor: true,
	}
}

func (nb Notebook) GetType() string {
	return nb.Type
}

func (nb Notebook) PrintDetails() {
	fmt.Printf("%s: Ядра:[%d], Оперативная память: [%d], Наличие монитора: [%v]\n", nb.Type, nb.Core, nb.Memory, nb.Monitor)
}

func (nb Notebook) String() string {
	return fmt.Sprintf("%s: Ядра:[%d], Оперативная память: [%d], Наличие монитора: [%v]\n", nb.Type, nb.Core, nb.Memory, nb.Monitor)
}
