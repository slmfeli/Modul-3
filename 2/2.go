package main
import ("fmt")

func f(x int) int {
	return x * x
}

func g(x int) int {
	return x - 2
}

func h(x int) int {
	return x + 1
}

func fogoh(x int) int {
	return f(g(h(x)))
}

func gohof(x int) int {
	return g(h(f(x)))
}

func hofog(x int) int {
	return h(f(g(x)))
}

func main() {
	var a, b, c int

	fmt.Print("Masukkan 3 angka (a b c): ")
	fmt.Scan(&a, &b, &c)

	resultFogoh := fogoh(a)
	resultGohof := gohof(b)
	resultHofog := hofog(c)
	
	fmt.Println(resultFogoh)
	fmt.Println(resultGohof)
	fmt.Println(resultHofog)
}