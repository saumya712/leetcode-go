func majorityElement(nums []int) int {
    freq:=make(map[int]int)
    n:=len(nums)-1

    for _, num := range(nums){
        freq[num]++
        if freq[num]>n/2{
            return num
        }
    }

    return -1
}
