package valueobjects

type DishType int

const (
	generic DishType = iota + 1
	fish
	seafood
	meat
	vegetable
	pasta
	cereal
	legume
	soup
	snack
	dessert
)

func (dt DishType) String() string {
	return [...]string{"generic", "fish", "seafood", "meat", "pasta", "cereal", "legume", "soup", "snack", "dessert"}[dt-1]
}

func (dt DishType) EnumIndex() int {
	return int(dt)
}
