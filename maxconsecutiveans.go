func maxConsecutiveAnswers(answerkey string, k int) int {
    left:=0
    ans:=0
    maxfreq:=0
    freq:=make(map[byte]int)

    for right:=0;right<len(answerkey);right++ {
        freq[answerkey[right]]++

        if freq[answerkey[right]] > maxfreq {
            maxfreq=freq[answerkey[right]]
        }
        if (right-left+1)-maxfreq > k {
            freq[answerkey[left]]--
            left++
        }
        if right-left+1 > ans {
            ans=right-left+1
        }

    }
    return ans
}
