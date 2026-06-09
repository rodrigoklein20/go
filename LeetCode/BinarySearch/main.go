package main

import "log"

func main() {
	arr := []int{-2, -1, 0, 4, 5, 7, 9, 12, 15}

	if search(arr, 11) == -1 {
		log.Print("Não encontrado")
	} else {
		log.Print("Encontrado")
	}
}

//Retorna o índice onde encontrou o 'target'
func search(nums []int, target int) int {
	low := 0
	high := len(nums) - 1

	for low <= high {
		mid := low + (high-low)/2

		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			low = mid + 1
			continue
		} else if nums[mid] > target {
			high = mid - 1
			continue
		}
	}

	return -1
}
