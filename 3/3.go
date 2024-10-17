package main
import (
	"fmt"
	"math"
)

func jarak(a, b, c, d float64) float64 {
	return math.Sqrt(math.Pow(a-c, 2) + math.Pow(b-d, 2))
}

func diDalam(cx, cy, r, x, y float64) bool {
	return jarak(cx, cy, x, y) <= r
}

func main() {
	var cx1, cy1, r1, cx2, cy2, r2, x, y float64
	
	fmt.Print("Lingkaran 1 (cx, cy, radius): ")
	fmt.Scan(&cx1, &cy1, &r1)
	fmt.Print("Lingkaran 2 (cx, cy, radius): ")
	fmt.Scan(&cx2, &cy2, &r2)
	fmt.Print("Titik sembarang (x, y): ")
	fmt.Scan(&x, &y)

	insideLingkaran1 := diDalam(cx1, cy1, r1, x, y)
	insideLingkaran2 := diDalam(cx2, cy2, r2, x, y)

	if insideLingkaran1 && insideLingkaran2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if insideLingkaran1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if insideLingkaran2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}