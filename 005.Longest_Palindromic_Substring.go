func longestPalindrome(s string) string {
    slen := len(s)
    
    if slen <= 1 {
        return s
    }
    
    dp := make([][]int, slen)
    
    for i := 0; i < slen; i++ {
        a := make([]int,slen)
        dp[i] = a
        dp[i][i] = 1    
    }
    
    start, end := 0, 0
    maxLen := 1
    for j := 1; j < slen; j++ {    // j 
        for i := 0; i < j; i++ {     // i shift from left to right until j 
            if s[i] == s[j] {
                if j - i == 1 {
                    dp[i][j] = 2
                }else{
                    if dp[i+1][j-1] > 0 {
                        dp[i][j] = dp[i+1][j-1] + 2
                    }
                }
            }
            
            if dp[i][j] > maxLen {
                maxLen = dp[i][j]
                start = i
                end = j
            }
        }
    }
    
    return s[start:end+1]
    
}