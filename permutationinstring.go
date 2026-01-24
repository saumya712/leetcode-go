func checkInclusion(s1 string, s2 string) bool {
    if len(s1) > len(s2) {
        return false
    }

    s1Count := make(map[byte]int)
    windowCount := make(map[byte]int)

    // frequency of s1
    for i := 0; i < len(s1); i++ {
        s1Count[s1[i]]++
    }

    left := 0

    for right := 0; right < len(s2); right++ {
        // add current char to window
        windowCount[s2[right]]++

        // if window size exceeds s1 length, shrink from left
        if right-left+1 > len(s1) {
            windowCount[s2[left]]--
            if windowCount[s2[left]] == 0 {
                delete(windowCount, s2[left])
            }
            left++
        }

        // check if window matches s1
        if right-left+1 == len(s1) && mapsEqual(s1Count, windowCount) {
            return true
        }
    }

    return false
}

// helper to compare two maps
func mapsEqual(a, b map[byte]int) bool {
    if len(a) != len(b) {
        return false
    }
    for k, v := range a {
        if b[k] != v {
            return false
        }
    }
    return true
}
