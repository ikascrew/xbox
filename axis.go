package xbox

type Axis struct {
	Name  string
	index int
	Value int
}

func NewAxis(i int, n string, v int) *Axis {
	ax := Axis{}

	ax.index = i
	ax.Name = n
	ax.Value = v
	return &ax
}
