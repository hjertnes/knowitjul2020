package day1

import (
	"errors"
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func convertListOfStringsToListOfInts(numbers []string) []int{
	ints := make([]int, 0)

	for _, n := range numbers {
		i, _ := strconv.Atoi(n)
		ints = append(ints, i)
	}

	sort.Ints(ints)

	return ints
}

func readData() string{
	data, _ := ioutil.ReadFile("./day1/data.txt")
	return string(data)
}

func findMissingNumber(numbers []int) (int, error){
	for index, number := range numbers {
		if number-index != 1{
			return number-1, nil
		}
	}

	return -1, errors.New("no match found")

}

func Run(){
	fmt.Println("1st of December")
	fmt.Println("Task: The dataset contains all the numbers between 1 and 100 000, except one. Find the missing one.")

	data := readData()
	listOfStrings := strings.Split(data, ",")
	numbers := convertListOfStringsToListOfInts(listOfStrings)
	match, err := findMissingNumber(numbers)
	if err != nil{
		fmt.Println("No match found")
		return
	}

	fmt.Println(fmt.Sprintf("Answer: %v", match))
}
