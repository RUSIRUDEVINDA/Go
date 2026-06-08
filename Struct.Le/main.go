package main

import "fmt"

type Color struct {
	Name string
	Hex  int
}

func slicesToObjects(colorNames []string, hexValues []int) []Color {
	var colors []Color

	for i := 0; i < len(colorNames); i++ {
		colors = append(colors, Color{
			Name: colorNames[i],
			Hex:  hexValues[i],
		})
	}
	return colors
}

func main() {
	names := []string{"Red", "Green", "Blue"}
	hexes := []int{0xFF0000, 0x00FF00, 0x0000FF}

	colors := slicesToObjects(names, hexes)

	fmt.Println(colors)
}