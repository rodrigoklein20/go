package main

// 14. Longest Common Prefix
// Write a function to find the longest common prefix string amongst an array of strings.

// If there is no common prefix, return an empty string "".

// Example 1:

// Input: strs = ["flower","flow","flight"]
// Output: "fl"
// Example 2:

// Input: strs = ["dog","racecar","car"]
// Output: ""
// Explanation: There is no common prefix among the input strings.

// Constraints:

// 1 <= strs.length <= 200
// 0 <= strs[i].length <= 200
// strs[i] consists of only lowercase English letters if it is non-empty.

import (
	"log"
	"strings"
)

func main() {
	arr := []string{"flower", "flow", "flight"}
	log.Printf("Esperado: fl  retorno: %s", longestCommonPrefix(arr))

	arr = []string{"flower", "flower", "flower"}
	log.Printf("Esperado: flower  retorno: %s", longestCommonPrefix(arr))

	arr = []string{"dog", "racecar", "car"}
	log.Printf("Esperado:   retorno: %s", longestCommonPrefix(arr))

	arr = []string{""}
	log.Printf("Esperado:   retorno: %s", longestCommonPrefix(arr))

	arr = []string{"a"}
	log.Printf("Esperado: a  retorno: %s", longestCommonPrefix(arr))
}

func longestCommonPrefix(strs []string) string {
	var retorno string
	var prefixo string

	if len(strs) == 1 {
		return strs[0]
	}

	processa := true
	indice := 1
	for processa {
		retorno = prefixo
		prefixo = ""
		for _, palavra := range strs {
			if prefixo != "" {
				if !strings.HasPrefix(palavra, prefixo) {
					retorno = prefixo[0 : len(prefixo)-1]
					processa = false
					break
				}
				continue
			}

			if indice > len(palavra) {
				processa = false
				break
			}
			prefixo = palavra[0:indice]
		}
		indice++
	}
	return retorno
}
