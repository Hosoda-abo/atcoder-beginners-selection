package main

import "fmt"

func main() {
	var N, count int
	fmt.Scan(&N)

	A := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&A[i])
	}

	flag := true
	count = 0

	for flag {
		for i := 0; i < N; i++ {
			if A[i]%2 == 0 {
				A[i] /= 2
			} else {
				flag = false
				break
			}
		}
		if flag {
			count++
		}
	}

	fmt.Println(count)
}
