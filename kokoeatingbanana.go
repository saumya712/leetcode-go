func minEatingSpeed(piles []int, h int) int {
    left:=1
    right:=0
    for _,p:=range piles {
        if p > right {
            right=p
        }
    }
    ans:=right

    for right>=left {
        mid:=(left+right)/2
        hrs:=0

        for _,p := range piles {
            hrs+=(p+mid-1)/mid
        }

        if hrs>h {
            left=mid+1
        } else {
            ans=mid
            right=mid-1
        }
    }
    return ans
}
