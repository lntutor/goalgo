package algorithm

import (
	"fmt"
)

func maxDraw() {
	var T int
	fmt.Scanf("%d", &T)
	for i := 0; i < T; i++ {
		var input int
		fmt.Scanf("%d", &input)
		fmt.Println(input + 1)
	}
}
