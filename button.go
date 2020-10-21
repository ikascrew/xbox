package xbox

type Button struct {
	Name  string
	index int
}

func NewButton(i int, n string) *Button {
	b := Button{}
	b.index = i
	b.Name = n
	return &b
}
