package XMLParser

import (
	"encoding/xml"
	"strings"
)

// Menu defines the root element of the XML file.
type Menu struct {
	Food []Food `xml:"food"`
}

// Food contains the definition of the job. For instance it contains:
// - Name
// - Price
// - Calories
type Food struct {
	Name     string `xml:"name"`
	Price    string `xml:"price"`
	Calories string `xml:"calories"`
}

// ToString returns the stringify representation of the Menu
func (menu *Menu) ToString() string {
	foods := make([]string, len(menu.Food))
	for index, element := range menu.Food {
		foods[index] = element.ToString()
	}
	return strings.Join(foods, ", ")
}

// ToString returns the stringify representation of the Menu
func (food *Food) ToString() string {
	return food.Name + " price $" + food.Price
}

// Parser unmarshal them XMl payload received by
func Parser(in []byte) (menu Menu) {
	xml.Unmarshal(in, &menu)
	return menu
}
