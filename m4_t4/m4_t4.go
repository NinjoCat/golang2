/*
Задача №4. Шахматная доска
Вход: размер шахматной доски, от 0 до 20
Выход: вывести на экран эту доску, заполняя поля нулями и единицами

Пример:
Вход: 5
Выход:
    0 1 0 1 0
    1 0 1 0 1
    0 1 0 1 0
    1 0 1 0 1
    0 1 0 1 0
*/

package main

import (
	"fmt"
)

func main() {
	var size int
	fmt.Println("Введите размер доски")
	fmt.Scanf("%d", &size)

	f := 1
	for i := 0; i < size; i++ {
		f = getF(f)
		fn := f
		for n := 0; n < size; n++ {
			fn = getF(fn)
			fmt.Printf("%d ", fn)
		}
		fmt.Println("")
	}
}

func getF(f int) int {
	if f == 1 {
		return 0
	} else {
		return 1
	}
}
