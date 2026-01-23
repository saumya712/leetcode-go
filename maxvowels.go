func maxVowels(s string, k int) int {
    isvowel:=func (c byte) bool {
        return c=='a'||c=='e'||c=='i'||c=='o'||c=='u'
    }

    currentvowel:=0
    maxvowel:=0

    for i:=0;i<k;i++ {
        if isvowel(s[i]) {
            currentvowel++
        }
    }
    maxvowel=currentvowel
    
    for right:=k;right<len(s);right++ {
        left:=right-k

        if isvowel(s[left]) {
            currentvowel--
        }

        if isvowel(s[right]) {
            currentvowel++
        }

        if currentvowel>maxvowel {
            maxvowel=currentvowel
        }
    }
    return maxvowel
}
