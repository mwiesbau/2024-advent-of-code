package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
)

func main() {

	// inputList1 := []int {3, 4, 2, 1, 3, 3}
	// inputList2 := []int {4, 3, 5, 3, 9, 3}

	inputList1, inputList2, err := LoadInput("input.txt")

	if err != nil {
		log.Printf("Failed to load file: %v", err)
	}

	if len(inputList1) != len(inputList2) {
		log.Fatal("The input lists are not equal, check the logic!")
	}

	log.Printf("Parsed input to lists of length: %v", len(inputList1))
	sort.Ints(inputList1)
	sort.Ints(inputList2)

	log.Printf("Completed sorting lists")
	totalDistance := ComputeDistane(inputList1, inputList2)

	log.Printf("Total distance: %v\n", totalDistance)

	log.Printf("The similartity score is: %v", computeSimilarityScore(inputList1, inputList2))
}

func computeSimilarityScore(list1 []int, list2 []int) int {

	var totalScore = 0

	m := make(map[int]int)

	// initialize the map
	for _, number := range list2 {
		m[number] = m[number] + 1
	}

	// compute scores using the map
	for _, number := range list1 {
		totalScore += number * m[number]
	}
	return totalScore
}

func ComputeDistane(list1 []int, list2 []int) int {
	var totalDistance = 0
	for i := 0; i < len(list1); i++ {
		delta := math.Abs(float64(list1[i] - list2[i]))
		totalDistance += int(delta)
	}

	return totalDistance
}

func LoadInput(filePath string) ([]int, []int, error) {
	log.Printf("loading input from: %v", filePath)
	file, err := os.Open(filePath)
	if err != nil {
		return []int{}, []int{}, err
	}

	fileReader := bufio.NewScanner(file)
	fileReader.Split(bufio.ScanWords)

	var fileLines []string

	for fileReader.Scan() {
		fileLines = append(fileLines, fileReader.Text())
	}

	file.Close()

	var list1 []int
	var list2 []int

	for i := 0; i < len(fileLines); i++ {
		number, err := strconv.Atoi(fileLines[i])

		if err != nil {
			return []int{}, []int{}, err
		}

		if i%2 == 0 {
			list2 = append(list2, number)
		} else {
			list1 = append(list1, number)
		}
	}

	return list1, list2, err
}
