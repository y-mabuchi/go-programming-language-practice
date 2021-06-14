package weightconv

import "fmt"

type Kilogram float64
type Pond float64

func (k Kilogram) String() string {
	return fmt.Sprintf("%gkg", k)
}

func (p Pond) String() string {
	return fmt.Sprintf("%glb", p)
}
