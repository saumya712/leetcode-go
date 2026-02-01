func singleNumber(nums []int) int {
    result := 0

    for bit := 0; bit < 32; bit++ {
        count := 0

        for _, num := range nums {
            if (num>>bit)&1 == 1 {
                count++
            }
        }

        if count%3 != 0 {
            result |= (1 << bit)
        }
    }

    return result
}
