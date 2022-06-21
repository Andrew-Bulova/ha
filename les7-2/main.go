package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input := "1 9 3 4 -5"

	var (
		intInput  []int
		minMaxArr []string
		result    string
	)

	for _, v := range strings.Split(input, " ") {
		intV, _ := strconv.Atoi(v)
		intInput = append(intInput, intV)
	}
	sort.Ints(intInput)
	minMaxArr = append(minMaxArr, strconv.Itoa(intInput[len(intInput)-1]))
	minMaxArr = append(minMaxArr, strconv.Itoa(intInput[0]))

	result = strings.Join(minMaxArr, " ")

	fmt.Println("З поданого ряду чисел -", input)
	fmt.Println("найбільше та найменше числа відповідно - ", result)

}
