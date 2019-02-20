package q1

func twoSum(nums []int, target int) (result []int) {
	numMap := make(map[int]int)

	for idx2, num := range nums {
		if idx1, ok := numMap[target-num]; ok {
			result = []int{idx1, idx2}
			break
		}
		numMap[num] = idx2
	}
	return
}
