class Solution:
    def sortArray(self, nums: List[int]) -> List[int]:
        

        def merge_sort(nums):
            if len(nums) <= 1:
                return nums
            mid=len(nums)//2
            left=merge_sort(nums[:mid])
            right=merge_sort(nums[mid:])
            return merge(left,right)
        def merge(left,right):
            result=[]
            i,j=0,0
            while i < len(left) and j < len(right):
                if left[i] > right[j]:
                    result.append(right[j])
                    j+=1
                else:
                    result.append(left[i])
                    i+=1
            result.extend(left[i:]) # append remaining elements
            result.extend(right[j:])
            return result
        return merge_sort(nums)