func maxArea(height []int) int {
    left, right:=0,len(height)-1

    maxarea:=0

    for left < right {
        h:=min(height[left],height[right])
        w:=right-left
        area:=w*h

        if area>maxarea {
            maxarea=area
        }
        if height[left]>height[right] {
            right--
        } else {
            left++
        }
    }
    return maxarea
}

func min(a,b int) int{
    if a<b {
        return a
    }
    return b
}
