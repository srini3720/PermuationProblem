package main

import (
	"errors"
	"fmt"
	"time"
	"runtime"
)

var studentSubsectionallotmentIndex int = 0
var subsectionSequence []int
var subsectionSequenceBlockIndex int
var permutationlimit int = 100


func main() {
	
	start := time.Now()
	var subsectionArraysize int64 = 7
	var totalStuInsamecourseAndSub int64 = 1
	fmt.Println("subsectionArraysize :", subsectionArraysize)
	fmt.Println("totalStuInsamecourseAndSub :", totalStuInsamecourseAndSub)
	var i int64
	fmt.Println("------------------------HeapAlgorithm----------------------------")
	for i = 0; i < totalStuInsamecourseAndSub; i++ {
		var subsection int = NewSubSectionAlgorithm(subsectionArraysize)
		fmt.Println("student ", i+1, " is in index of Heap Method subsection :", subsection)
	}
	fmt.Println("Time for 10 students allotment ",time.Since(start))
	runtime.GC()
	PrintMemUsage()
}

func NewSubSectionAlgorithm(subsectionArraysize int64) int {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	if subsectionArraysize >= 10 {
		errors.New("Subsection array size should be less than 10")
	}
	var subsectionArray []int
	var permutesArray [][]int
	var blockSize int
	if len(subsectionSequence) != 0 && studentSubsectionallotmentIndex < len(subsectionSequence) {
		var randomSubsection = subsectionSequence[studentSubsectionallotmentIndex]
		studentSubsectionallotmentIndex++
		return randomSubsection
	}
	if len(subsectionSequence) >= studentSubsectionallotmentIndex {
		studentSubsectionallotmentIndex = 0
		fmt.Println("PermutationLimit :", permutationlimit)
		subsectionArray = makeArray(int(subsectionArraysize))
		permutesArray = permutationBlock(subsectionArray)
		fmt.Println("permutation Block  :", permutesArray)
		blockSize = len(permutesArray)
		fmt.Println("permutation Block size  :", blockSize)
		subsectionSequence = permutesArray[subsectionSequenceBlockIndex]
		fmt.Println("subsection sequence", subsectionSequence)
		var selectedSubsection = subsectionSequence[studentSubsectionallotmentIndex]
		subsectionSequenceBlockIndex++
		studentSubsectionallotmentIndex++
		return selectedSubsection

		// studentSubsectionallotmentIndex = studentSubsectionallotmentIndex + 1
	}
	return 0
}

func makeArray(num int) []int {
	arrayOne := make([]int, num)
	for i := 0; i < num; i++ {
		arrayOne[i] = i + 1
	}
	arrayTwo := make([]int, num)
	for i := 0; i < num; i++ {
		arrayTwo[i] = i + 1
	}
	var array []int = append(arrayOne, arrayTwo...)
	fmt.Println("subsection Array :", array)
	return array
}

func permutationBlock(array []int) (permutesArray [][]int) {
	var permutaionNumber int
	var permutaion func([]int, int)
	permutaion = func(permutes []int, size int) {
		if size == len(permutes) {
			if permutaionNumber == permutationlimit {
				return
			}
			permutesArray = append(permutesArray, append([]int{}, permutes...))
			permutaionNumber++
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


func bToMb(b uint64) uint64 {
    return b / 1024 
}

func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v KB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v KB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys Memory Usage = %v KB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}