package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var subsectionArraysize int64 = 4
	var totalStuInsamecourseAndSub int64 = 10
	fmt.Println("subsectionArraysize :", subsectionArraysize)
	fmt.Println("totalStuInsamecourseAndSub :", totalStuInsamecourseAndSub)
	var i int64
	for i = 0; i < totalStuInsamecourseAndSub; i++ {
		var index int64 = oldsubsectionalgorithm(subsectionArraysize, i)
		fmt.Println("student ", i+1, " is in index of (RoundRobinAlgo) :", index, "& subsection :", index+1)
	}
	fmt.Println("New Subsection Algorithm (Heap algorithm)-----------")
	NewSubSectionAlgorithm(subsectionArraysize, totalStuInsamecourseAndSub)

}
func oldsubsectionalgorithm(subsectionArraysize int64, totalStuInsamecourseAndSub int64) int64 {
	var SubsectionIDIndex int64
	SubsectionIDIndex = totalStuInsamecourseAndSub % subsectionArraysize
	return SubsectionIDIndex
}

func NewSubSectionAlgorithm(subsectionArraysize int64, totalStuInsamecourseAndSub int64) {
	if subsectionArraysize >= 10 {
		errors.New("Subsection array size should be less than 10")
		return
	}
	var subsectionArray = makeArray(int(subsectionArraysize))
	fmt.Println("subsectionArray :", subsectionArray)
	var permutesArray [][]int = permutationBlock(subsectionArray)
	fmt.Println("permutation Block  :", permutesArray)
	//random

	var blockSize int = len(permutesArray)
	var randomIndex int = randomNumber(blockSize)
	var randomSequence []int = permutesArray[randomIndex]
	fmt.Println("Random Number :", randomIndex)
	fmt.Println("Random sequence :", randomSequence)
	for i := 0; i < len(randomSequence); i++ {
		var index int = randomSequence[i]
		fmt.Println("student ", i+1, " is in index of (HeapAlgorithm) Subsection:", index)
	}
	randomSequence = []int{}
	randomIndex = randomNumber(blockSize)
	fmt.Println("Random Number :", randomIndex)
	randomSequence = permutesArray[randomIndex]
	fmt.Println("Random sequence :", randomSequence)
	for i := 0; i < len(randomSequence); i++ {
		var index int = randomSequence[i]
		fmt.Println("student ", i+1, " is in index of (HeapAlgorithm) Subsection:", index)
	}
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
