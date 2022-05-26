package main

import (
	"fmt"
	"time"
	"runtime"
	"math/rand"
	"container/list"
)


var subsectionSequence []int
var subsectionIndex int
var subsectionlist = list.New()

func main() {
	start := time.Now()
	var subsectionArraysize int64 = 10
	var totalStuInsamecourseAndSub int64 = 100
	fmt.Println("subsectionArraysize :", subsectionArraysize)
	fmt.Println("totalStuInsamecourseAndSub :", totalStuInsamecourseAndSub)
	fmt.Println("------------------------DoubleRandamization----------------------------")
	
	
	var i int64
	for i = 0; i < totalStuInsamecourseAndSub; i++ {
		var subsection int = doublerandamization(subsectionArraysize)
		fmt.Println("student ", i+1, " is in index of Double randamization subsection :", subsection)
	}

	for e := subsectionlist.Front(); e != nil; e = e.Next() {
		fmt.Println("subsection List", e.Value)
	}
	fmt.Println("Time for ",totalStuInsamecourseAndSub, "students allotment ",time.Since(start))
	runtime.GC()
	PrintMemUsage()
}



func doublerandamization(subsectionArraysize int64) int {
	
	if (len(subsectionSequence) != 0 && subsectionIndex != 0 && subsectionIndex < len(subsectionSequence)) {
		var randomSubsection int =subsectionSequence[subsectionIndex]
		subsectionIndex = subsectionIndex + 1
		return randomSubsection
	}
	if len(subsectionSequence) >= subsectionIndex {
		subsectionIndex =0
		var subsectionArray []int= makeArray(int(subsectionArraysize))
		fmt.Println("subsection Array :", subsectionArray)
		var shuffledArray []int = shuffleArray(subsectionArray)
		subsectionSequence = shuffledArray
		fmt.Println("shuffled Array :", shuffledArray)
		var allotment int = allotmentStudents(shuffledArray)
		subsectionlist.PushBack(subsectionArray)
		
		return allotment
	}
	
	// fmt.Println("student is in index of Heap Method subsection :", allotment)
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
	return array
}
func shuffleArray(array []int) []int {
	for i:=0; i<len(array); i++ {
		var randomNumber int = randomNumber(len(array))
		var temp int= array[i] 
		array[i] = array[randomNumber]
		array[randomNumber] = temp
	}
	// if belongsToList(array) {
	// 	return shuffleArray(array)
	// }
	for e := subsectionlist.Front(); e != nil; e = e.Next() {

		fmt.Println("subsection List", e.Value)
	}
	return array
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

func randomNumber(number int) int {
	var x1 = rand.NewSource(time.Now().UnixNano())
	var y1 = rand.New(x1)
	var randomIndex int = y1.Intn(number)
	return randomIndex

}

func allotmentStudents(array []int ) int {
	subsectionIndex = subsectionIndex + 1
	var randomNumber int = randomNumber(len(array))
	return array[randomNumber]

}

// func belongsToList(lookup string) bool {
// 	for _, val := range subsectionlist {
// 		if val == lookup {
// 			return true
// 		}
// 	}
// 	return false
// }