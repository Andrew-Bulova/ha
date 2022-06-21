package main

import (
	"fmt"
)

func main() {
	arr := []int{4, 1, 4, -4, 6, 3, 8, 8}
	var newArr, result []int

	for i := len(arr) - 1; i >= 0; i-- {
		var isDuplicated bool = false
		for k := i - 1; k >= 0; k-- {
			if arr[i] == arr[k] {
				isDuplicated = true
				break
			}
		}

		if !isDuplicated {
			newArr = append(newArr, arr[i])
		}

	}

	for i := len(newArr) - 1; i >= 0; i-- {
		result = append(result, newArr[i])
	}
	fmt.Println(result)

	/*
	   Приклад даних на вхід: [3, 4, 4, 3, 6, 3]
	   видаляємо 3 по індексу 0
	   видаляємо 4 по індексу 1
	   видаляємо 3 по індексу 3
	   Правильний результат: [3, 4, 6]

	   !!! В наведеному результаті порядок елементів змінений.
	   В умові не зазначено, що масив повинен сортуватися, тому
	   вивід програми буде [4 6 3]
	*/
}
