package main

import (
	"fmt"
	"time"
	"runtime"
	"math/rand"
	"container/list"
	"strings"
)


var subsectionSequence []int
var subsectionlist = list.New()

func main() {
	start := time.Now()
	var subsectionArraysize int64 = 2
	var totalStuInsamecourseAndSub int64 = 1
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



// func doublerandamization(subsectionArraysize int64) int {
	
// 	if (len(subsectionSequence) != 0 && subsectionIndex != 0 && subsectionIndex < len(subsectionSequence)) {
// 		var randomSubsection int =subsectionSequence[subsectionIndex]
// 		subsectionIndex = subsectionIndex + 1
// 		return randomSubsection
// 	}
// 	if len(subsectionSequence) >= subsectionIndex {
// 		subsectionIndex =0
// 		var subsectionArray []int= makeArray(int(subsectionArraysize))
// 		fmt.Println("subsection Array :", subsectionArray)
// 		var shuffledArray []int = shuffleArray(subsectionArray)
// 		subsectionSequence = shuffledArray
// 		fmt.Println("shuffled Array :", shuffledArray)
// 		var allotment int = allotmentStudents(shuffledArray)
// 		var arrayString string = arrayToString(shuffledArray, ",")
// 		subsectionlist.PushBack(arrayString)
		
// 		return allotment
// 	}
// 	return 0
// }
func doublerandamization(subsectionArraysize int64) int {
	
	if (len(subsectionSequence) == 0) {
		var subsectionArray []int= makeArray(int(subsectionArraysize))
		var shuffledArray []int = shuffleArray(subsectionArray)
		subsectionSequence = shuffledArray
		fmt.Println("shuffled Array :", shuffledArray)
		var allotment int = allotmentStudents(shuffledArray)
		var arrayString string = arrayToString(shuffledArray, ",")
		subsectionlist.PushBack(arrayString)
		
		return allotment
	}
	if len(subsectionSequence) != 0 {
		var allotment int = allotmentStudents(subsectionSequence)
		return allotment
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
	return array
}
func shuffleArray(array []int) []int {
	for i:=0; i<len(array); i++ {
		var randomNumber int = randomNumber(len(array))
		var temp int= array[i] 
		array[i] = array[randomNumber]
		array[randomNumber] = temp
	}
	var arrayStr string = arrayToString(array, ",")

	for e := subsectionlist.Front(); e != nil; e = e.Next() {
		if e.Value == arrayStr {
			return shuffleArray(array)
		}
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

func randomNumber(number int) int {
	var x1 = rand.NewSource(time.Now().UnixNano())
	var y1 = rand.New(x1)
	var randomIndex int = y1.Intn(number)
	return randomIndex

}

func allotmentStudents(array []int ) int {
	
	var randomNumber int = randomNumber(len(array))
	subsectionSequence = removeIndexFromArray(array,randomNumber)
	fmt.Println("subsectionSequenceAfter removed index :",subsectionSequence)
	return array[randomNumber]

}


func arrayToString(array []int, delim string) string {
    return strings.Trim(strings.Replace(fmt.Sprint(array), " ", delim, -1), "[]")
    //return strings.Trim(strings.Join(strings.Split(fmt.Sprint(array), " "), delim), "[]")
    //return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(array)), delim), "[]")
}

func removeIndexFromArray(array []int, Index int) []int {
    return append(array[:Index], array[Index+1:]...)
}