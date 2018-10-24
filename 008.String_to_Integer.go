func myAtoi(str string) int {
    const INT32_MAX = 2147483647  //int(^uint(0) >> 1)
    const INT32_MIN = -2147483648 //^INT_MAX
    
    slen := len(str)
    
    spaceFilter := false
    isNegative := false
    isPositive := false
    hasFirstNum := false
    
    ret := 0
    
    for i:=0; i<slen; i++ {
        if str[i] == ' ' && !spaceFilter {
            continue
        }
        
        spaceFilter = true 
        
        if str[i] == '-' && !isNegative && !hasFirstNum {
            isNegative = true
            continue
        }
        
        if str[i] == '+' && !isPositive && !hasFirstNum {
            isPositive = true
            continue
        }
        
        if isNegative && isPositive {  //'  +-1232' or '-+1232'
            return 0
        }
        
        if str[i] >= 48 && str[i] <=57 {    //num 0-9 ascii 
            hasFirstNum = true
            ret = ret*10 + int(str[i]) - 48
            
            if isNegative {
                if -1*ret < INT32_MIN {
                    return INT32_MIN 
                }
            }else{
                if ret > INT32_MAX {
                    return INT32_MAX
                }
            }
            
        } else {
            break
        }
    }
    
    if isNegative {
        ret *= -1
    }
    
    return ret
}