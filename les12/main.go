package main

import (
	"fmt"
	"sync"
)

// Конкурентно порахувати суму кожного слайсу int, та роздрукувати результат.
// Потрібно використовувати WaitGroup.
// Приклад:
// [ [ 4, 6 ], [ 7, 9 ] ]
// Результат друку:
// Порядок друку не важливий.
// “slice 1: 10”
// “slice 2: 16”
func main() {
	var wg sync.WaitGroup
	n := [][]int{
		{2, 6, 9, 24},
		{7, 3, 94, 3, 0},
		{4, 2, 8, 35},
	}

	wg.Add(len(n))
	for key, value := range n {
		go func(key int, value []int) {
			defer wg.Done()
			fmt.Printf("slice %v: %v\n", key, sum(value))
		}(key, value)
	}

	wg.Wait()
}

func sum(arr []int) int {
	res := 0
	for _, value := range arr {
		res += value
	}

	return res

}
