package arrays_and_slices

//import "fmt"

func Sum(arr []int) int {
	result := 0
	for _, v := range arr {
		result += v
	}
	return result
}

func SumAll(numsToSum ...[]int) []int {
	var result []int
	for _, v := range numsToSum {
		result = append(result, Sum(v))
	}
	return result
}

func SumAllTails(numsToSum ...[]int) []int {
	var result []int
	for _, v := range numsToSum {
		if len(v) == 0 {
			result = append(result, 0)
		} else {
			result = append(result, Sum(v[1:]))
		}
	}
	return result
}
