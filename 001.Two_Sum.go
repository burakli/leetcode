/***
func twoSum(nums []int, target int) []int {
    lenNums := len(nums)
    
    for i := 0; i < lenNums; i++ {
        for j := i + 1; j < lenNums; j++ {
            if nums[i] + nums[j] == target {
                return []int{i, j}
            }
        }
    }
    
    return []int{}
        
}
**/

/**
func twoSum(nums []int, target int) []int {
    lenNums := len(nums)
    
    numsMap := make(map[int]int, lenNums)
    
    for i := 0; i < lenNums; i++ {
        numsMap[nums[i]] = i       
    } 
    
    for i := 0; i < lenNums; i++ {
        numB := target - nums[i]
        if k,ok := numsMap[numB]; ok && k != i {
            return []int{i,k}
        }
    }
    
    return []int{}
        
}
**/

func twoSum(nums []int, target int) []int {
    lenNums := len(nums)
    
    numsMap := make(map[int]int, lenNums)
    
    for i := 0; i < lenNums; i++ {
        numB := target - nums[i]
        if k,ok := numsMap[numB]; ok && k != i {
            return []int{k,i}    //k before i
        }
        numsMap[nums[i]] = i
    }
    
    return []int{}
    
}
