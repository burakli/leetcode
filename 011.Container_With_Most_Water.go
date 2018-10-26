//Two Pointer Approach, Time complexity: O(N)
func maxArea(height []int) int {

    maxArea := 0
    i, j := 0, len(height)-1
    
    for {
        if (i == j) {
            break
        }
        
        h := height[i]
        if height[i] > height[j] {
            h = height[j]      //h = min of height[i],height[j]
        }
        area := (j - i) * h
        if area > maxArea {
            maxArea = area
        }
        
        if height[i] < height[j] {
            i++
        }else{
            j--
        }
    }
    
    return maxArea
}


/**
// Brute Force, Time complexity : O(n^2)
func maxArea(height []int) int {

    maxArea := 0
    
    hlen := len(height)
    
    for i := 0; i < hlen; i++ {    
        for j := i+1; j < hlen; j++ {
            h := height[i]
            if height[i] > height[j] {
                h = height[j]      //h = min of height[i],height[j]
            }
            area := (j - i) * h
            if area > maxArea {
                maxArea = area
            }
        }
    }
    
    return maxArea
}
*/
