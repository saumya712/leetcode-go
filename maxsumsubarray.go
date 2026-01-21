func maxSubArray(nums []int) int {
    maxsum:=nums[0]

    currentsum:=0

    for i:=0;i<len(nums);i++ {
        currentsum+=nums[i]

        if maxsum<currentsum {
            maxsum=currentsum
        }

        if currentsum<0 {
            currentsum=0
        }

    }

    return maxsum


}
