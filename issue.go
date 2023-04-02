func sortArray(nums []int) []int {
    if len(nums) < 2 {
        return nums
    }
    sort.Ints(nums)
   return nums
}

class Solution:
    def sortArray(self, nums: List[int]) -> List[int]:
        nums.sort()
        return nums
