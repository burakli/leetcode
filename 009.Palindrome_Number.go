func isPalindrome(x int) bool {
    
    if x < 0 {
        return false
    }

    if x/10 == 0 {
        return true
    }
    
    t := x
    xlen := 0
    b := 1
    for {
        xlen++
        if t/10 == 0 {
            break
        }
        t = t/10
        b *= 10
        
    }

    
    for i := 0; i < xlen/2; i++ {
        if x % 10 != x / b {
            return false
        }
        
        x = x % b 
        x = x / 10
        b = b/100
    }
    
    return true
}
