package string_utils

import "strings"

func CompareStrings(str1 string, str2 string) bool {
	str1 = strings.ToLower(str1)
	str2 = strings.ToLower(str2)

	str1 = strings.TrimSpace(str1)
	str2 = strings.TrimSpace(str2)

	str1 = strings.Replace(str1, " ", "-", -1)
	str2 = strings.Replace(str2, " ", "-", -1)

	str1Slice := strings.Split(str1, "-")
	str2Slice := strings.Split(str2, "-")

	lenSlice1 := 0
	for _, slice1 := range str1Slice {
		slice1 = strings.TrimSpace(slice1)
		if slice1 != "" {
			lenSlice1 += 1
		}
	}
	lenSlice2 := 0
	for _, slice2 := range str2Slice {
		slice2 = strings.TrimSpace(slice2)
		if slice2 != "" {
			lenSlice2 += 1
		}
	}
	if lenSlice1 != lenSlice2 {
		return false
	}

	for _, slice1 := range str1Slice {
		foundSlice := false
		slice1 = strings.TrimSpace(slice1)
		if slice1 != "" {
			for _, slice2 := range str2Slice {
				slice2 = strings.TrimSpace(slice2)
				if slice2 != "" {
					if slice1 == slice2 {
						foundSlice = true
						break
					}
				}
			}
			if !foundSlice {
				return false
			}
		}
	}
	return true
}
