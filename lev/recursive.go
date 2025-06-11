package lev

import (
	"strings"
)

func Recursive(string1 string, string2 string) int {
	distance := 0

	if len(string2) > len(string1) {
		temp := string1
		string1 = string2
		string2 = temp
	}

	if len(string1) == 0 || len(string2) == 0 {
		distance = max(len(string1), len(string2))
		return distance
	}

	splitString1 := strings.Split(string1, "")
	splitString2 := strings.Split(string2, "")

	for len(splitString1) > 0 && len(splitString2) > 0 {
		if splitString1[len(splitString1)-1] == splitString2[len(splitString2)-1] {
			splitString1 = splitString1[:len(splitString1)-1]
			splitString2 = splitString2[:len(splitString2)-1]
		} else {
			break
		}
	}

	for len(splitString1) > 0 && len(splitString2) > 0 {

		if splitString1[0] == splitString2[0] {
			splitString1 = splitString1[1:]
			splitString2 = splitString2[1:]
		} else {
			break
		}
	}

	if len(splitString1) == 0 || len(splitString2) == 0 {
		distance = max(len(splitString1), len(splitString2))
		return distance
	}

	distance = min(
		Recursive(strings.Join(splitString1[1:], ""), strings.Join(splitString2, ""))+1,
		Recursive(strings.Join(splitString1, ""), strings.Join(splitString2[1:], ""))+1,
		Recursive(strings.Join(splitString1[1:], ""), strings.Join(splitString2[1:], ""))+1)

	return distance
}
