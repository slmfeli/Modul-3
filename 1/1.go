package main
import ("fmt")

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

func permutation(n, r int) int {
	return factorial(n) / factorial(n-r)
}

func combination(n, r int) int {
	return factorial(n) / (factorial(r) * factorial(n-r))
}

func main() {
	var a, b, c, d int

	fmt.Print("Masukkan 4 angka (a b c d): ")
	fmt.Scan(&a, &b, &c, &d)

	permutationC := permutation(a, c)
	combinationC := combination(a, c)
	permutationD := permutation(b, d)
	combinationD := combination(b, d)

	fmt.Println(permutationC, combinationC)
	fmt.Println(permutationD, combinationD)
}