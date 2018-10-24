func reverse(x int) int {
    const INT32_MAX = 2147483647  //int(^uint(0) >> 1)
    const INT32_MIN = -2147483648 //^INT_MAX
    
    if x > INT32_MAX ||  x < INT32_MIN {
        return 0
    }
    
    ret := 0
    t := x
    
    if t < 0 {
        t *= -1    
    }
    
    for {
        ret = ret*10 + t%10
        if ret > INT32_MAX ||  ret < INT32_MIN {
            return 0
        }
        t = t/10 
        if t == 0 {
            break
        }
    }
    
    if x < 0 {
        ret *= -1
    }
    
    return ret
}

/***
func reverse(x int) int {
    var buffer bytes.Buffer
    
    xStr := strconv.Itoa(x)
    
    if xStr[0] == '-' {
        buffer.Write([]byte{'-'})
    }
    
    xlen := len(xStr)
    
    last := xlen-1
    for i:= xlen-1; i>=0; i-- {
        if xStr[i] != '0' {
            last = i
            break
        }
    }
    
    for i:=last; i>=0; i-- {
        if xStr[i] != '-' {
            buffer.Write([]byte{xStr[i]})
        }
    }
    
    result, _ :=  strconv.Atoi(buffer.String())
    
    return result
  
}**/