/*
Написать функцию.
Входные аргументы функции: количество бутылок от 0 до 200.
Функция должна вернуть количество и слово бутыл<?> с правильным окончанием.
Пример :
In: 5
Out: 5 бутылок

In: 1
Out: 1 бутылка

In: 22
Out: 22 бутылки
*/

package main

import (
	"fmt"
	"golang_homework2/bottles.com/botlib"
	"log"
)

func main() {
	var bottlesNumber int

	fmt.Println("Введите количество бутылок от 0 до 200:")
	_, err := fmt.Scanf("%d\n", &bottlesNumber)

	if err != nil {
		log.Fatal("Epic fail: " + err.Error())
	}

	str, err := botlib.PrintBottle(bottlesNumber)

	if err != nil {
		log.Fatal("Epic fail: " + err.Error())
	}

	fmt.Println(str)
}
