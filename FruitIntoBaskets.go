func totalFruit(fruits []int) int {
    left:=0
    fruitcount:=make(map[int]int)
    maxfruit:=0

    for right:=0;right<len(fruits);right++ {
        fruitcount[fruits[right]]++

        for len(fruitcount)>2 {
            fruitcount[fruits[left]]--
            if fruitcount[fruits[left]]==0 {
                delete(fruitcount,fruits[left])
            }
            left++
        }
        currentlen:=right-left+1
        if currentlen>maxfruit {
           maxfruit=currentlen
        }
    }
    return maxfruit
}
