func sortArray(nums []int) []int {
    if len(nums) < 2 {
        return nums
    }
    sort.Ints(nums)
   return nums
}
