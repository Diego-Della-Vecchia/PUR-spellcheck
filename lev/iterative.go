package lev

import (
	"strings"
)

func Iterative(string1 string, string2 string) int {
	distance := 0

	if len(string2) > len(string1) {
		temp := string1
		string1 = string2
		string2 = temp
	}

	if len(string2) == 0 {
		distance = max(len(string1), len(string2))
		return distance
	}

	splitString1 := strings.Split(string1, "")
	splitString2 := strings.Split(string2, "")

	for len(splitString2) > 0 {
		if splitString1[len(splitString1)-1] == splitString2[len(splitString2)-1] {
			splitString1 = splitString1[:len(splitString1)-1]
			splitString2 = splitString2[:len(splitString2)-1]
		} else {
			break
		}
	}

	for len(splitString2) > 0 {

		if splitString1[0] == splitString2[0] {
			splitString1 = splitString1[1:]
			splitString2 = splitString2[1:]
		} else {
			break
		}
	}

	if len(splitString2) == 0 {
		distance = max(len(splitString1), len(splitString2))
		return distance
	}

	matrix := initMatrix(len(splitString1)+1, len(splitString2)+1)

	for i := 1; i <= len(splitString1); i++ {
		for j := 1; j <= len(splitString2); j++ {
			substitutionCost := 0
			if splitString1[i-1] != splitString2[j-1] {

				substitutionCost = 1
			}
			matrix[i][j] = min(
				matrix[i-1][j]+1,                  // Deletion
				matrix[i][j-1]+1,                  // Insertion
				matrix[i-1][j-1]+substitutionCost, // Substitution
			)
		}
	}

	distance = matrix[len(splitString1)][len(splitString2)]

	return distance

}

func initMatrix(len1 int, len2 int) [][]int {
	matrix := make([][]int, len1)
	for i := range matrix {
		matrix[i] = make([]int, len2)
	}

	for i := range matrix {
		matrix[i][0] = i
	}
	for j := range matrix[0] {
		matrix[0][j] = j
	}

	return matrix
}
