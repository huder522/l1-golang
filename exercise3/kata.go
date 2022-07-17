package main

import "fmt"

func TwoSum(numbers []int, target int) [2]int {
  
  var resultArray [2]int
  
  for i := 0; i < len(numbers)-1; i++ {
		for j := i+1; j < len(numbers); j++ {
			if numbers[i] + numbers[j] == target {
				resultArray[0] = i
        resultArray[1] = j
			}
		}
	}
	return resultArray
}

func main() {

	fmt.Println(TwoSum([]int{1, 2, 3}, 4))
	fmt.Println(TwoSum([]int{1234, 5678, 9012}, 14690))
	fmt.Println(TwoSum([]int{2, 2, 3}, 4))

}