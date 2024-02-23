package main

func addUpper(n int) int {
	res := 0
	for i := 1; i <= n; i++ {
		res += i
	}
	return res
}

func sub(n1 int, n2 int) int {
	res := n1 + n2
	return res
}
