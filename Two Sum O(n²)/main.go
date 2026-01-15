package main

import "fmt"

func twoSum(nums []int, target int) []int {
    result := []int{}
    for i := 0; i < len(nums); i++ {
        for j := 0 + 1; j < len(nums); j++ {
            operation := nums[i] + nums[j]
            if operation == target {
                result = append(result, i, j) 
                return result
            }
        }
    }
    result = append(result, -1)
    return result
}

func main() {
    nums := []int{3, 5, 7, 2}
    target := 8


    ts := twoSum(nums, target)
    if ts[0] == -1 {
        fmt.Printf("Impossível concluir encontrar o target neste array %d \n", nums)
        return
    } 
    fmt.Printf("Os indíces: %d com os valores: %d e %d resultam no target: %d", ts, nums[ts[0]], nums[ts[1]], target)
}