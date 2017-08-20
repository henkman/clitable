package clitable

import (
	"fmt"
	"os"
	"testing"
)

type City struct {
	Name       string
	Country    string
	Population int
}

func (c City) GetValue(i int) string {
	switch i {
	case 0:
		return c.Name
	case 1:
		return c.Country
	case 2:
		return fmt.Sprint(c.Population)
	}
	panic("invalid index")
}

func TestWrite(t *testing.T) {
	cities := []Row{
		City{
			Name:       "London",
			Country:    "UK",
			Population: 8700000,
		},
		City{
			Name:       "Berlin",
			Country:    "Germany",
			Population: 3500000,
		},
		City{
			Name:       "Paris",
			Country:    "France",
			Population: 2200000,
		},
		City{
			Name:       "Saint-Étienne",
			Country:    "France",
			Population: 171260,
		},
	}
	Write(os.Stdout,
		[]string{"Name", "Country", "Population"},
		cities)
}
