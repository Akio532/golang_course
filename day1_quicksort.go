package main

func main() {
	println((SolutionBinaryGap(1456))) //место для ввода числа
}

func SolutionBinaryGap(N int) int {
	n := N
	var num int = 0
	var count2 int = 0
	var count int = 0
	for n > 0 {
		if n%2 == 0 {
			count++
			if count > num {
				num = count
			}
		} else {
			count = 0
		}

		n = n / 2
		count2++
	}
	return num
}
