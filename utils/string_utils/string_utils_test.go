package string_utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCompareStrings2CompareStringsSameOrder(t *testing.T)  {
	str1 := "abc def GHI-jkl"
	str2 := "ABc DEF GHI-jkl"
	result := CompareStrings(str1, str2)
	assert.EqualValues(t, result, true)
}

func TestCompareStrings2CompareStringsDifferentOrder(t *testing.T)  {
	str1 := "abc def GHI-jkl"
	str2 := " DEF -jkl ABc-GHI"
	result := CompareStrings(str1, str2)
	assert.EqualValues(t, result, true)
}

func TestCompareStrings2CompareStringsEmptySlices(t *testing.T)  {
	str1 := "abc def GHI-jkl"
	str2 := " DEF -jkl ABc-GHI  - - - -   -"
	result := CompareStrings(str1, str2)
	assert.EqualValues(t, result, true)
}

func TestCompareStrings2CompareStringsEmptySlicesChangedOrder(t *testing.T)  {
	str2 := "abc def GHI-jkl"
	str1 := " DEF -jkl ABc-GHI  - - - -   -"
	result := CompareStrings(str1, str2)
	assert.EqualValues(t, result, true)
}

func TestCompareStrings2CompareStringsDifferentLength(t *testing.T)  {
	str1 := "abc def GHI-jkl"
	str2 := " DEF -jkl ABc-GHI  - - - - xxx"
	result := CompareStrings(str1, str2)
	assert.EqualValues(t, result, false)
}

func TestCompareStrings2CompareStringsDifferentLengthChangedOrder1(t *testing.T)  {
	str2 := "abc def GHI-jkl"
	str1 := " DEF -jkl ABc-GHI  - - - - xxx"
	result := CompareStrings(str1, str2)
	assert.EqualValues(t, result, false)
}

func TestCompareStrings2CompareStringsDifferentLengthChangedOrder2(t *testing.T)  {
	str2 := "xxx--abc def GHI-jkl"
	str1 := " DEF -jkl ABc-GHI  - - - - xxx"
	result := CompareStrings(str1, str2)
	assert.EqualValues(t, result, true)
}
