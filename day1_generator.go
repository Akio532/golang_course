package main

func main() {
	var start int = 20
	var n int = 20
	arr := make([]int, n)
	arr = SolutionSquareGenerator(start, n)
	for i := 0; i < n; i++ {
		println(arr[i])
	}

}

func SolutionSquareGenerator(start int, n int) []int {
	squares := make([]int, n)

	for i := 0; i < n; i++ {
		squares[i] = start * start
		start++
	}
	return squares
}
