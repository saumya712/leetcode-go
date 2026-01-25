func findMaxAverage(nums []int, k int) float64 {
    
    sum:=0
    
    for i:=0;i<k;i++ {
        sum+=nums[i]
    }
    maxsum:=sum

    for i:=k;i<len(nums);i++ {
        sum=sum-nums[i-k]+nums[i]
        if maxsum<sum {
            maxsum=sum
        }

    }
    return float64(maxsum)/float64(k)
}
