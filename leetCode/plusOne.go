package main
import "fmt"

func plusOne(digits []int) []int {
    for i:= len(digits)-1; i >= 0; i--{
        if digits[i] < 9{
            digits[i]++
            return digits
        }
        digits[i] = 0
    }
    return append([]int{1}, digits...)
}

func main(){
	result := plusOne([]int{9,9,9})
	fmt.Println(result)
}