func characterReplacement(s string, k int) int {
    freq:=make(map[byte]int)

    left:=0
    maxfreq:=0
    ans:=0

    for right:=0;right<len(s);right++ {
        freq[s[right]]++

        if freq[s[right]]>maxfreq {
            maxfreq=freq[s[right]]
        }

        if right-left+1 - maxfreq >k {
            freq[s[left]]--
            left++
        }

        if right-left+1 > ans {
            ans=right-left+1
        }
    }

    return ans
}
