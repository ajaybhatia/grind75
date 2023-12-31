package solutions

func TwoSum(nums []int, target int) []int {
	output := make([]int, 2)
	for i, v1 := range nums {
		for j, v2 := range nums {
			if i != j && v1+v2 == target {
				output[0] = i
				output[1] = j
				return output
			}
		}
	}
	return output
}
