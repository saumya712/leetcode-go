func numberOfSubarrays(nums []int, k int) int {

    return atmost(nums,k)-atmost(nums,k-1)
    
}
func atmost(nums []int,k  int) int{
    left:=0
    oddcount:=0
    count:=0

    for right:=0;right<len(nums);right++ {
        if nums[right]%2==1 {
            oddcount++
        }

        for oddcount > k {
            if nums[left]%2==1 {
                oddcount--
            }
            left++
        }

        
        count += right-left+1
        
    }

    return count
}
