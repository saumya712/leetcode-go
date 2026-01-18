func twoSum(num []int, target int) []int {
    left,right:=0,len(num)-1

    for right>left {
        sum := num[right]+num[left]

        if sum==target {
            return []int{left+1,right+1}
        }else if sum>target{
            right--
        }else{
            left++
        }
    }

    return nil
}
