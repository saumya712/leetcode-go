func canConstruct(ransomNote string, magazine string) bool {
    seen:=make(map[byte]int)

    for i:=0;i<len(magazine);i++ {
        seen[magazine[i]]++
    }

    for j:=0;j<len(ransomNote);j++ {
        if seen[ransomNote[j]]>0 {
            seen[ransomNote[j]]--
        } else if seen[ransomNote[j]]==0 {
            return false
        }
        
    }
    return true
}
