func findTheDifference(s string, t string) byte {
    seen:=make(map[byte]int)

    for i:=0;i<len(s);i++ {
        seen[s[i]]++
    }
    for j:=0;j<len(t);j++ {
        seen[t[j]]--
        if seen[t[j]]<0 {
            return t[j]
        }
    }
    return 0
}
    
