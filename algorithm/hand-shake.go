package algorithm

import "fmt"

func solve(n int) int {
	return n * (n - 1) / 2
}

func handshake() {
	var T int
	fmt.Scanf("%d", &T)
	for i := 0; i < T; i++ {
		var n int
		fmt.Scanf("%d", &n)
		fmt.Println(solve(n))
	}
}
