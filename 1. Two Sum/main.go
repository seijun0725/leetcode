func twoSum(nums []int, target int) []int {
    numbers := make(map[int]int)
    for i, num := range nums {
        if number, ok := numbers[num]; ok {
            return []int{number, i}
        }
        numbers[target-num] = i
    }
    return []int{}
}
