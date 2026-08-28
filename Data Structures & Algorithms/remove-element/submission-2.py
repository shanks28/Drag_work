class Solution:
    def removeElement(self, nums: List[int], val: int) -> int:
        ptr1=0
        ptr2=0

        while ptr1 < len(nums) and ptr2 < len(nums):
            if nums[ptr1]!=val:# not the element
                ptr1+=1
                ptr2+=1
            elif ptr2 < len(nums) and nums[ptr1]==val:
                while ptr2 < len(nums) and nums[ptr2]==val :
                    ptr2+=1
                if ptr1 < len(nums) and ptr2 < len(nums):
                    nums[ptr1],nums[ptr2]=nums[ptr2],nums[ptr1]
        return len(nums)-nums.count(val)

