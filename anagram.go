func findAnagrams(s string, p string) []int {
    freqp:=make(map[byte]int)

    left:=0
    indexes:=[]int{}

    for i:=0;i<len(p);i++ {
        freqp[p[i]]++
    }
    
    freqw:=make(map[byte]int)
    
    for right:=0;right<len(s);right++ {
        freqw[s[right]]++

        if right-left+1 > len(p) {
            freqw[s[left]]--
            if freqw[s[left]]==0 {
                delete(freqw,s[left])
            }
            left++
        }
        if right-left+1 == len(p) {
            if mapsEqual(freqw, freqp) {
                indexes=append(indexes, left)
            }
        }
    }
    return indexes
}
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
