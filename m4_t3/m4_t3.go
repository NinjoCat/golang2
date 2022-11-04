/*
Написать 3 функции.
Даны координаты трех точек(x1, y1, x2, y2, x3, y3), значения(целые) которых >= 0.
Первая функция проверяет, что можно построить треугольник по заданным точкам
Вторая функция вычисляет площадь треугольника.
Третья функция должна определить, является ли треугольник прямоугольным.
*/
package main

import (
	"fmt"
	"golang_homework2/triangles.com/triangleslib"
	"log"
)

func main() {
	var x1, y1, x2, y2, x3, y3 int
	fmt.Println("Введите координаты точек в формате: x1,y1,x2,y2,x3,y3: ")
	_, err := fmt.Scanf("%d,%d,%d,%d,%d,%d", &x1, &y1, &x2, &y2, &x3, &y3)

	if err != nil {
		log.Fatal("Epic fail: " + err.Error())
	}

	//1. Можно ли построить треугольник.
	if triangleslib.IsTriangle(x1, y1, x2, y2, x3, y3) {
		fmt.Println("Треугольник построить можно")
	} else {
		fmt.Println("Треугольник построить нельзя")
		return
	}

	//2. Площадь треугольника.
	fmt.Println("Площадь треугольника = ", triangleslib.Square(x1, y1, x2, y2, x3, y3))

	//3. Является ли треугольник прямоугольный.
	if triangleslib.IsRightTriangle(x1, y1, x2, y2, x3, y3) {
		fmt.Println("Треугольник прямоугольный")
	} else {
		fmt.Println("Треугольник непрямоугольный")
		return
	}
}
