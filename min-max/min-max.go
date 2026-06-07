package main

import "fmt"
import "bufio"
import "os"
import "strings"
import "strconv"

func ReadNumbers() []int {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Введите числа: ")
	scanner.Scan()
	line := scanner.Text()
	fmt.Println("Вы ввели: ", line)

	parts := strings.Split(line, " ")
	numbers := make([]int, 0, len(parts))

	for _, part := range parts {
		number, error := strconv.Atoi(part)
		if error != nil {
			fmt.Println("Ошибка преобразования: ", part, error)	
		}
		numbers = append(numbers, number)
	}
	return numbers
}

func FindMinNumber(numbers []int) int {
	minNum := numbers[0]
	for _, number := range numbers {
		if number < minNum {
			minNum = number
		}
	}
	return minNum
}

func FindMaxNumber(numbers []int) int {
	maxNum := numbers[0]
	for _, number := range numbers {
		if number > maxNum {
			maxNum = number
		}
	}
	return maxNum
}

func FindAvgNumber(numbers []int) int {
	sumNum := 0
	for _, number := range numbers {
		sumNum += number
	}
	return sumNum / len(numbers)
}

func calculate() {
	numbers := ReadNumbers()
	if len(numbers) == 0 {
		fmt.Println("Массив числел пустой. Завершение программы")
		return	
	}	
	
	fmt.Println("Среднее значение: ", FindAvgNumber(numbers))
	fmt.Println("Максимальное значение: ", FindMaxNumber(numbers))
	fmt.Println("Минимальное значение: ", FindMinNumber(numbers))
}


func main() {
	calculate()	
}
