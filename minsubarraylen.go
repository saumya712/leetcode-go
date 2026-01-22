func minSubArrayLen(target int, nums []int) int {
    left:=0
    sum:=0
    minlen:=len(nums)+1

    for right:=0;right<len(nums);right++ {
        sum+=nums[right]

        for sum>=target {
            currentlen:=right-left+1
            if currentlen<minlen {
                minlen=currentlen
            }
            sum-=nums[left]
            left++
        }
    }

    if minlen==len(nums)+1 {
        return 0
    }
    return minlen
}
