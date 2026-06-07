package main

import "fmt"

func Reverse(nums []int) []int {
	if len(nums) == 0 {
		return nums	
	}

	leftPoint := 0
	rightPoint := len(nums) - 1
	
	for leftPoint < rightPoint {
		temp := nums[leftPoint]
		nums[leftPoint] = nums[rightPoint]
		nums[rightPoint] = temp		
		leftPoint++
		rightPoint--
	}
	return nums
}

func Deduplicate(nums []int) []int {
	set := make(map[int]bool)
	var deduplicated []int = make([]int, 0, len(nums))
	for _, num := range nums {
		if set[num] == false {
			deduplicated = append(deduplicated, num)
			set[num] = true
		}
	}
	return deduplicated
}

func Chunk(nums []int, chunkSize int) [][]int {
	if chunkSize <= 0 {
		fmt.Println("Размер чанка должен быть больше 0.")
		return [][]int{}
	}
	
	if len(nums) == 0 {
		fmt.Println("Пустой набор чисел")
		return [][]int{}
	}
	
	var chunkCount int = len(nums) / chunkSize
	
	if len(nums) % chunkSize > 0 {
		chunkCount++
	}

	// fmt.Println("Количество чанков: ", chunkCount)

	chunks := make([][]int, chunkCount)

	for i := 0; i < chunkCount; i++ {
		chunks[i] = make([]int, 0, chunkSize)
	}

	for index, num := range nums {
		//fmt.Println("Число: ", num)
		//fmt.Println("Индекс:", index)
		//fmt.Println("Чанк:", index / chunkSize)
		//fmt.Println("Индекс в чанке: ", index % chunkSize)
		//fmt.Println()
		chunks[index / chunkSize] = append(chunks[index / chunkSize], num)
	}
	
	return chunks
}


func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 8, 9, 9}
	fmt.Println(Deduplicate(Reverse(nums)))
	fmt.Println(Chunk(nums, 3))
}













