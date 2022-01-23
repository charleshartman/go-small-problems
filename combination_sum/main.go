package main

func combinationSum(candidates []int, target int) [][]int {
	var results [][]int
	var candidate []int
	backtrack(candidates, candidate, &results, target)
	return results
}

func backtrack(list []int, candidate []int, results *[][]int, target int) {
	if sum(candidate) == target {
		copyCandidate := make([]int, len(candidate))
		copy(copyCandidate, candidate)
		*results = append(*results, copyCandidate)
		return
	}

	for i := 0; i < len(list); i++ {
		num := list[i]

		if num+sum(candidate) > target {
			continue
		}

		candidate = append(candidate, list[i])          // take
		backtrack(list[i:], candidate, results, target) // explore
		candidate = candidate[:len(candidate)-1]        // clean up
	}
}

func sum(s []int) (result int) {
	for _, v := range s {
		result += v
	}
	return
}
