package main

import (
	"log"
	"slices"
)

func main() {

	log.Printf("Quantidade restante: %d", removeElement([]int{3, 2, 2, 3}, 3))

	log.Printf("Quantidade restante: %d", removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))
}

func removeElement(nums []int, val int) int {
	for {
		indice := slices.Index(nums, val)
		if indice == -1 {
			return len(nums)
		}
		nums = slices.Delete(nums, indice, indice+1)
	}
}
