package arrays

func Sum(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func SumAll(groups ...[]int) []int {
	sums := make([]int, len(groups))
	for i, group := range groups {
		sums[i] = Sum(group)
	}
	return sums
}

func SumAllTails(groups ...[]int) []int {
	sums := make([]int, len(groups))
	for i, group := range groups {
		if len(group) == 0 {
			sums[i] = 0
		} else {
			sums[i] = Sum(group[1:])
		}
	}
	return sums
}
