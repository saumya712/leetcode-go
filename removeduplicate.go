func removeDuplicates(nums []int) int {
    if len(nums)==0 {
        return 0
    }
    unik:=1

    for i:=1;i<len(nums);i++ {
        if nums[i]!=nums[i-1] {
            nums[unik]=nums[i]
            unik++
        }

    }
    return unik
}
