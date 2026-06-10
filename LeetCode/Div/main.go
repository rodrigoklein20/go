package main

import (
	"log"
	"slices"
)

func main() {
	// log.Printf("RESULTADO: %v", shuffle([]int{2, 5, 1, 3, 4, 7}, 3))

	// log.Printf("Máximo de números 1 consecutivos %d", findMaxConsecutiveOnes([]int{1, 1, 0, 1, 1, 1, 1, 0, 1, 1}))

	// findErrorNums([]int{1, 2, 2, 4, 5})
	// findErrorNums([]int{1, 1})
	// findErrorNums([]int{2, 2})
	// findErrorNums([]int{4, 2, 2, 1})

	smallerNumbersThanCurrent([]int{8, 2, 9, 3, 4})
}

func shuffle(nums []int, n int) []int {

	var resultado []int
	for i := 0; i < n; i++ {
		resultado = append(resultado, nums[i])
		resultado = append(resultado, nums[i+n])
	}
	return resultado
}

func findMaxConsecutiveOnes(nums []int) int {
	var max int

	var tst int
	for _, n := range nums {
		if n == 1 {
			tst += 1
			if tst > max {
				max = tst
			}
		} else if tst > max {
			max = tst
			tst = 0
		} else {
			tst = 0
		}
	}
	return max
}

func findDuplicate(nums []int) int {
	for i, n := range nums {
		for j := 0; j < len(nums); j++ {
			if i == j {
				continue
			} else if n == nums[j] {
				return n
			}
		}
	}
	return 0
}

func findMissing(nums []int) int {
	end := len(nums)

	for i := 1; i <= end; i++ {
		if !slices.Contains(nums, i) {
			return i
		}
	}
	return 0
}

func findErrorNums(nums []int) []int {
	log.Print("Entrada ", nums)

	ret := []int{}

	ret = append(ret, findDuplicate(nums))
	ret = append(ret, findMissing(nums))

	log.Print("Saída", ret)
	return ret
}

func smallerNumbersThanCurrent(nums []int) []int {
	var maior int
	var ocurr int

	for _, n := range nums {
		if n > maior {
			maior = n
			ocurr = 1
		} else if n == maior {
			ocurr++
		}
	}
	log.Printf("%d menores", len(nums)-ocurr)

	return []int{}
}
