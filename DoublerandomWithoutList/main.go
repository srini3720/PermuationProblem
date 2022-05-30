package main

import (
	"fmt"
	"time"
	"runtime"
	"math/rand"
)



func main() {
	start := time.Now()
	var subsectionArraysize int64 = 10
	var totalStuInsamecourseAndSub int64 = 100
	fmt.Println("subsectionArraysize :", subsectionArraysize)
	fmt.Println("totalStuInsamecourseAndSub :", totalStuInsamecourseAndSub)
	fmt.Println("------------------------DoubleRandamization----------------------------")
	
	
	var i int64
	for i = 0; i < totalStuInsamecourseAndSub; i++ {
		var subsection int64 = doublerandamization(subsectionArraysize)
		fmt.Println("student ", i+1, " is in index of Double randamization subsection :", subsection)
	}

	fmt.Println("Time for ",totalStuInsamecourseAndSub, "students allotment ",time.Since(start))
	runtime.GC()
	PrintMemUsage()
}




func doublerandamization(subsectionArraysize int64) int64 {
	var subsectionArray []int64= makeArrayAndShuffleForTwoN(subsectionArraysize)
	fmt.Println("shuffled Array :", subsectionArray)
	var randomNumber int64 = randomNumber(int64(len(subsectionArray)))
	fmt.Println("randomNumber :", randomNumber)
	var allotedSubsection int64 = subsectionArray[randomNumber]
	return allotedSubsection - 1
}
func makeArrayAndShuffleForTwoN(num int64) []int64 {
	var i int64
	arrayOne := make([]int64, num)
	for i = 0; i < num; i++ {
		arrayOne[i] = i + 1
	}
	arrayTwo := make([]int64, num)
	for i = 0; i < num; i++ {
		arrayTwo[i] = i + 1
	}
	var array []int64 = append(arrayOne, arrayTwo...)
	for i:=0; i<len(array); i++ {

		var randomNumber int64 = rand.Int63n(int64(len(array)))
		var temp int64= array[i] 
		array[i] = array[randomNumber]
		array[randomNumber] = temp
	}
	return array
}

func bToMb(b uint64) uint64 {
    return b / 1024 
}

func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v KB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v KB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys Memory Usage = %v KB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func randomNumber(number int64) int64 {
	// var NewRand = rand.New(rand.NewSource(time.Now().UnixNano()))
	var randomIndex int64 = rand.Int63n(number)
	return randomIndex

}

