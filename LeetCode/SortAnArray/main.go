package main

import (
	"log"
	"math/rand"
	"slices"
	"time"
)

// 912. Sort an Array
// Given an array of integers nums, sort the array in ascending order and return it.
// You must solve the problem without using any built-in functions in O(nlog(n)) time complexity and with the smallest space complexity
const tam = 5000

func main() {
	nums := make([]int, tam)

	for i := range nums {
		nums[i] = rand.Intn(tam)
	}
	sortArrayRuim(nums)

	for i := range nums {
		nums[i] = rand.Intn(tam)
	}
	sortWithPack(nums)
}

func sortWithPack(nums []int) []int {
	start := time.Now()
	slices.Sort(nums)
	log.Printf("Duração sortWithPack: %f", time.Since(start).Seconds())
	return nums
}

func sortArrayRuim(nums []int) []int {
	start := time.Now()
	ret := []int{}
	i := 0
	for {
		if len(nums) == 0 {
			break
		}
		maior := false

		for j := 0; j < len(nums); j++ {
			if nums[i] > nums[j] {
				maior = true
			}
		}

		if !maior {
			ret = append(ret, nums[i])
			nums = slices.Delete(nums, i, i+1)
			i = 0
		} else {
			i++
		}
	}
	log.Printf("Duração sortArrayRuim: %f", time.Since(start).Seconds())
	return ret
}
