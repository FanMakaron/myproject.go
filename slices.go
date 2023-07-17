// https://go.dev/tour/moretypes/18
package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	result := [][]uint8{}

	for i := 0; i < dy; i++ {
		var res []uint8
		for j := 0; j < dx; j++ {
			res = append(res, 0)
		}
		result = append(result, res)
	}

	return result
}

func main() {
	pic.Show(Pic)
}
