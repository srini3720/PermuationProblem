package main

import (
	"fmt"
)

func permutationBlock(array []int) (permutesArray [][]int) {
	var permutaionNumber int;
	var permutaion func([]int, int)
	permutaion =func (permutes []int, size int) {
		if size == len(permutes) {
			if permutaionNumber ==permutationlimit{
				return
			}
			permutesArray = append(permutesArray, append([]int{}, permutes...))
			permutaionNumber++
			fmt.Println("permutationlimit :",permutationlimit, "permutaionNumber :",permutaionNumber)
		} else {
			for i:=size; i<len(array); i++ {
				permutes[size],permutes[i] = permutes[i],permutes[size]
				permutaion(permutes, size+1)
				permutes[size],permutes[i] = permutes[i],permutes[size]
			}
		}
	}
	permutaion(array,0)
	return permutesArray
} 


func makeArray(num int) []int {
	array := make([]int, num)
	for i := 0; i < num; i++ {
		array[i] = i + 1
	}
	return array
}
var  permutationlimit int =30
func main() {
	num :=6
	array := makeArray(num)
	fmt.Println("array func  :",array)
	fmt.Println("permutation func  :",permutationBlock(array))
	
}