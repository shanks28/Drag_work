class Solution:
    def rotate(self, nums: List[int], k: int) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        rotated=[]
        to_rotate=k%len(nums)
        for i in range(to_rotate):
            n=nums.pop(-1)
            rotated.append(n)
        rotated.reverse()
        rotated.extend(nums)
        nums[:]=rotated