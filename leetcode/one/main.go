package main

func main() {
	nums := []int{1, 2, 3, 4}
	target := 3
	twoSum(nums, target)
	twoSum1(nums, target)
}

func twoSum(nums []int, target int) []int {
	//遍历数组，获取每个下标和对应值
	//判断值其中两个值相加等于target
	//获取对应的两个值的下标
	//var result []int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == target-nums[j] {
				//result = append(result, i, j)
				return []int{i, j}
			}
		}
	}
	return nil
}

func twoSum1(nums []int, target int) []int {
	rest := []int{}
	m := make(map[int]int)
	for k, v := range nums {
		//
		val := target - v
		_, ok := m[val]
		if ok {
			rest = []int{m[val], k}
		}
		m[v] = k
	}
	return rest
}
