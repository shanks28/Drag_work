class Solution:
    def productExceptSelf(self, nums: List[int]) -> List[int]:
        left_product=[1]*len(nums)
        right_product=[1]*len(nums)
        cur_product=1
        i=1
        left_product[0]=1
        while i < len(nums):
            cur_product*=nums[i-1]
            left_product[i]=cur_product
            i+=1
        i=len(nums)-2
        right_product[len(nums)-1]=1
        cur_product=1
        while i >=0:
            cur_product*=nums[i+1]
            right_product[i]=cur_product
            i-=1
        i=0
        res=[0]*len(nums)
        while i < len(right_product):
            res[i]=left_product[i]*right_product[i]
            i+=1
        return res