func isPalindrome(s string) bool {
    i,j:=0,len(s)-1

    for i<j {

    

        if !isalpha(s[i]){
          i++
          continue
        }

        if !isalpha(s[j]){
          j--
          continue
        }

        if tolower(s[j]) != tolower(s[i]){
          return false 
        }
        i++
        j--
    }
return true
}

func isalpha(c byte) bool {
    return (c>='a'&& c<='z') ||
           (c>='A'&& c<='Z') ||
           (c>='0'&& c<='9')
}

func tolower(c byte) byte {
    if c>='A'&& c<='Z' {
        return c + ('a' - 'A')
    }
    return c

}
