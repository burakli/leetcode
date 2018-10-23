func lengthOfLongestSubstring(s string) int {
    
    smap := make(map[byte]int)
    
    cur := 0
    maxlen := 0
    
    for i:=0; i<len(s); i++ {
        
        if last, ok := smap[s[i]]; ok {
            if last >= cur {
                cur = last + 1   //Move to the next non repeat position
            }
        }
        
        if i - cur + 1 > maxlen {
            maxlen = i - cur + 1
        }
        
        smap[s[i]] = i
    }
    
    return maxlen
}


