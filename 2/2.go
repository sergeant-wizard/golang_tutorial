package main

import "golang.org/x/tour/pic"

func pixel(x, y int) uint8 {
	return uint8(x + y/2)
}

func Pic(dx, dy int) [][]uint8 {
	ret := make([][]uint8, dx)
	for x := range ret {
		ret[x] = make([]uint8, dy)
		for y := range ret[x] {
			ret[x][y] = pixel(x, y)
		}
	}
	return ret
}

func main() {
	pic.Show(Pic)
}
