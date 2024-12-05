package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


func main() {

  inputArray, err := LoadFile("input.txt")

  if err != nil {
    fmt.Printf("Error loading file: %v", err)
  }

  var safeCounter = 0
  for problemId, list := range inputArray {
    var isSafe = false
    var index = -1
    var newIndex = -1
    isSafe, index = reportIsSafe(list)
    fmt.Printf("[%v] %v is safe: %v\n", problemId, list, isSafe)
    if isSafe {
      safeCounter++
    } else {
      fmt.Printf("[%v] Issue at index %v %v, retrying\n", problemId, index, list)
      newList := removeIndex(list, index)
      isSafe, newIndex = reportIsSafe(newList)
      
      if index > 0 && !isSafe {
        listRemovePrior := removeIndex(list, index-1)
        isSafe, _ = reportIsSafe(listRemovePrior)
        fmt.Printf("[%v] Removed prior and is safe: %v with list: %v\n", problemId, isSafe, listRemovePrior)
      }

      if index < len(list) && !isSafe {
        listRemoveNext := removeIndex(list, index+1)
        isSafe, _ = reportIsSafe(listRemoveNext)
        fmt.Printf("[%v] Removed next and is safe: %v with list: %v\n", problemId, isSafe, listRemoveNext)
      }

      if isSafe {
        safeCounter++
        fmt.Printf("[%v] %v is safe on second try\n", problemId, newList)
      } else {
        fmt.Printf("[%v] Still unsave, old: %v  %v at index: %v\n", problemId, list, newList, newIndex)
      }
    }
  }

  fmt.Printf("\n------\nFound %v safe reports\n", safeCounter)
}

func removeIndex(list []int, index int) []int {

  newList := []int{}
  newList = append(newList, list[:index]...)
  newList = append(newList, list[index+1:]...)
  return newList

}

func reportIsSafe(list []int) (bool, int) {
  var MAX_DELTA = 3
  //var IS_INCREASING = false

  //if list[0] < list[1] {
  //  IS_INCREASING = true
  //}

  var increaseCount = 0
  var decreaseCount = 0

  for i := 0; i < len(list) - 1; i++ {
  
    delta := list[i] - list[i+1]
    if delta > 0 {
      increaseCount++
    } 

    if delta < 0 {
      decreaseCount++
    }

    if delta == 0 {
      return false, i
    }

    if increaseCount > 0 && decreaseCount > 0 {
      return false, i
    }
  

    if delta > MAX_DELTA || delta < -MAX_DELTA {
      return false, i
    }

  }
  return true, -1
}



func LoadFile(fileName string) ([][]int, error) {

  file, err := os.Open(fileName)

  if err != nil {
    return [][]int{}, err
  }

  fileReader := bufio.NewScanner(file)
  fileReader.Split(bufio.ScanLines)

  var fileLines []string

  for fileReader.Scan() {
    fileLines = append(fileLines, fileReader.Text())
  }


  file.Close()
  

  resultList := [][]int{}
  // Split input and reate 2d array
  for i := 0; i < len(fileLines); i++ {
    words := strings.Split(fileLines[i], " ")
    
    resultList = append(resultList, []int{})
    for j := 0; j < len(words); j++ {

      value, err := strconv.Atoi(words[j])

      if err != nil {
        return [][]int{}, err
      }

      resultList[i] = append(resultList[i], value)
    }
  }
  
  return resultList, nil
}
