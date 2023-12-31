package main

import (
	"fmt"

	"github.com/ajaybhatia/leetcode/solutions"
)

func main() {
	// 01_two_sum
	fmt.Println()
	fmt.Println(">>> 01_two_sum <<<")
	fmt.Println(solutions.TwoSum([]int{3, 2, 4}, 6))
	fmt.Println("------------------")

	// 02_add_two_numbers
	fmt.Println()
	fmt.Println(">>> 02_add_two_numbers <<<")
	fmt.Println(solutions.IsValid("()"))
	fmt.Println(solutions.IsValid("()[]{}"))
	fmt.Println(solutions.IsValid("(]"))
	fmt.Println("------------------")
}
