func subarraysWithKDistinct(nums []int, k int) int {
    return atmost(nums,k)-atmost(nums,k-1)
}
func atmost(nums []int,k int) int {
    left:=0
    count:=0
    freq:=make(map[int]int)

    for right:=0;right<len(nums);right++ {
        freq[nums[right]]++

        for len(freq) > k {
            freq[nums[left]]--
            if freq[nums[left]]==0 {
                delete(freq,nums[left])
            }
            left++
        }
        count+=right-left+1
    }
    return count
}
