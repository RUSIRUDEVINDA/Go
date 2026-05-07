package main

import "fmt"

func main(){

    var nums = [6]int{2,3,4,1,7,8}

    for i:=1;i<len(nums)-2;i++ {
        if nums[i]>nums[i+1] && nums[i]>nums[i-1] {
            fmt.Printf("peak value is " , nums[i])
        }
    }

    
}