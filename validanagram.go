func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    freq := make(map[byte]int)

    // count characters in s
    for i := 0; i < len(s); i++ {
        freq[s[i]]++
    }

    // cancel characters using t
    for i := 0; i < len(t); i++ {
        freq[t[i]]--
        if freq[t[i]] < 0 {
            return false
        }
    }

    return true
}
