package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solve(arrayA, arrayB []int) [2]int {
	result := [2]int{0, 0}
	for i := 0; i < len(arrayA); i++ {
		if arrayA[i] > arrayB[i] {
			result[0]++
		} else if arrayA[i] < arrayB[i] {
			result[1]++
		}
	}
	return result
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var arrayA []int
	var arrayB []int
	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		array := strings.Split(line, " ")
		for i := 0; i < len(array); i++ {
			if count == 0 {
				tmp, _ := strconv.Atoi(array[i])
				arrayA = append(arrayA, tmp)
			} else {
				tmp, _ := strconv.Atoi(array[i])
				arrayB = append(arrayB, tmp)
			}

		}
		count++
	}
	result := solve(arrayA, arrayB)
	fmt.Printf("%v %v\n", result[0], result[1])
}
