func smallestDivisor(nums []int, threshold int) int {
    left:=1
    right:=max(nums)
    ans:=0
    
    for right >= left {
        mid:=(left+right)/2
     
        
        if candiv(nums,mid,threshold) {
            ans=mid
            right=mid-1
        } else {
            left=mid+1
        }
    }
    return ans
}
func candiv(nums []int,mid int,threshold int) bool {
    sum:=0

    for _,p :=range nums {
        sum+=(p+mid-1)/mid
    } 
    return sum<=threshold
}
func max(nums []int) int {
    x:=0
    for _,p := range nums {
        if p > x {
            x=p
        }
    }
    return x
}
