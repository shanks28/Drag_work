
import (
    "slices"
)

func longestConsecutive(nums []int) int {
    if len(nums) == 0 {
        return 0
    }

    slices.Sort(nums)

    nums = slices.Compact(nums)

    res := 1
    window := 1

    for i := 1; i < len(nums); i++ {
        if nums[i]-nums[i-1] == 1 {
            window++
        } else {
            res = max(res, window)
            window = 1
        }
    }

    return max(res, window) // Final check for trailing sequences
}