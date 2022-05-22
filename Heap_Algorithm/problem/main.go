package main

import "fmt"

func main() {
	var subsectionArraysize int64 = 3
	var totalStuInsamecourseAndSub int64 = 10
	fmt.Println("subsectionArraysize :", subsectionArraysize)
	fmt.Println("totalStuInsamecourseAndSub :", totalStuInsamecourseAndSub)
	var i int64
	for i = 0; i < totalStuInsamecourseAndSub; i++ {
		fmt.Println("student ", i+1, " is in subsection in (oldalgo)", oldsubsectionalgorithm(subsectionArraysize, i))
	}
	fmt.Println("New Subsection Algorithm (Heap algorithm)-----------")
	NewSubSectionAlgorithm(subsectionArraysize, totalStuInsamecourseAndSub)

}
func oldsubsectionalgorithm(subsectionArraysize int64, totalStuInsamecourseAndSub int64) int64 {
	var SubsectionIDIndex int64
	SubsectionIDIndex = totalStuInsamecourseAndSub % subsectionArraysize
	return SubsectionIDIndex
}

func NewSubSectionAlgorithm(subsectionArraysize int64, totalStuInsamecourseAndSub int64)  {
	var subsectionArray = makeArray(int(subsectionArraysize))
	fmt.Println("subsectionArray :", subsectionArray)
	fmt.Println("permutation Block  :",permutationBlock(subsectionArray))
}


func makeArray(num int) []int {
	array := make([]int, num)
	for i := 0; i < num; i++ {
		array[i] = i + 1
	}
	return array
}


func permutationBlock(array []int) (permutesArray [][]int) {
	var permutaion func([]int, int)
	permutaion =func (permutes []int, size int) {
		if size == len(permutes) {
			permutesArray = append(permutesArray, append([]int{}, permutes...))
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