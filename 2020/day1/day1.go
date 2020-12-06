package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readInputIntoSlice(filepath string) []int {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var numbers []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, _ := strconv.Atoi(scanner.Text())
		numbers = append(numbers, number)
	}
	return numbers
}

func main() {
	numbers := readInputIntoSlice("input.txt")
	i, j := TwoSum(numbers, 2020)
	if i == -1 || j == -1 {
		log.Fatal("no sum found")
	}
	fmt.Printf("%d and %d sum up to 2020, their product is %d \n", numbers[i], numbers[j], numbers[i]*numbers[j])

}

// TwoSum returns the index of 2 integers in the slice numbers that sum up to target
func TwoSum(numbers []int, target int) (int, int) {
	ht := make(map[int]int)

	for i, num := range numbers {
		key := target - num
		if j, ok := ht[key]; ok {
			return j, i
		}
		ht[num] = i
	}
	return -1, -1
}