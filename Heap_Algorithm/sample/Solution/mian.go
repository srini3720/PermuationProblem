package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)
var arrayIndex int = 0
var randomSequence []int 

func main() {
	var subsectionArraysize int64 = 4
	var totalStuInsamecourseAndSub int64 = 20
	fmt.Println("subsectionArraysize :", subsectionArraysize)
	fmt.Println("totalStuInsamecourseAndSub :", totalStuInsamecourseAndSub)
	var i int64
	for i = 0; i < totalStuInsamecourseAndSub; i++ {
		var index int64 = oldsubsectionalgorithm(subsectionArraysize, i)
		fmt.Println("student ", i+1, " is in index of (RoundRobinAlgo) :", index, "& subsection :", index+1)
	}
	fmt.Println("------------------------Heap----------------------------")
	for i = 0; i < totalStuInsamecourseAndSub; i++ {
		var subsection int = NewSubSectionAlgorithm(subsectionArraysize)
		fmt.Println("student ", i+1, " is in index of Heap Method subsection :", subsection)
	}

}
func oldsubsectionalgorithm(subsectionArraysize int64, totalStuInsamecourseAndSub int64) int64 {
	var SubsectionIDIndex int64
	SubsectionIDIndex = totalStuInsamecourseAndSub % subsectionArraysize
	return SubsectionIDIndex
}

func NewSubSectionAlgorithm(subsectionArraysize int64) int {
	if subsectionArraysize >= 10 {
		errors.New("Subsection array size should be less than 10")
	}
	var subsectionArray []int
	var permutesArray [][]int
	var blockSize int
	var randomBlockNumber int

	if len(randomSequence) != 0 && arrayIndex < len(randomSequence) {
		var randomSubsection = randomSequence[arrayIndex]
		arrayIndex++
		return randomSubsection
	}  
	if len(randomSequence) >= arrayIndex {
		arrayIndex = 0
		subsectionArray = makeArray(int(subsectionArraysize))
		permutesArray = permutationBlock(subsectionArray)
		fmt.Println("permutation Block  :", permutesArray)
		blockSize = len(permutesArray)
		randomBlockNumber = randomNumber(int(blockSize))
		fmt.Println("Random  :", randomBlockNumber)
		randomSequence = permutesArray[randomBlockNumber]
		fmt.Println("Random sequence", randomSequence)
		var randomSubsection = randomSequence[arrayIndex]
		arrayIndex++
		return randomSubsection

		// arrayIndex = arrayIndex + 1
	} 
	return 0
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
	permutaion = func(permutes []int, size int) {
		if size == len(permutes) {
			permutesArray = append(permutesArray, append([]int{}, permutes...))
		} else {
			for i := size; i < len(array); i++ {
				permutes[size], permutes[i] = permutes[i], permutes[size]
				permutaion(permutes, size+1)
				permutes[size], permutes[i] = permutes[i], permutes[size]
			}
		}
	}
	permutaion(array, 0)
	return permutesArray
}

func randomNumber(number int) int {
	var x1 = rand.NewSource(time.Now().UnixNano())
	var y1 = rand.New(x1)
	var randomIndex int = y1.Intn(number)
	return randomIndex

}
