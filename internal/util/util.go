package util

import (
	"fmt"
	"math"
	"strconv"

	"github.com/samber/lo"
)

// FormatSf formats value with sf significant figures. The horizontal space used will remain
// static.
func FormatSf(value float64, sf int) string {
	digits := len(fmt.Sprint(int(value)))
	roundTo := sf - digits

	if roundTo < 1 {
		return fmt.Sprintf("%d", int(value))
	} else {
		return fmt.Sprintf(fmt.Sprintf("%%0.%df", roundTo), value)
	}
}

// FormatBytes formats bytes as gigabytes to sf significant figures.
func FormatBytes(bytes float64, sf int) string {
	return FormatSf(bytes/math.Pow(1024, 3), sf)
}

// AdjustColour adjusts a colour's brightness or darkness.
//
// Assumes colour begins with '#' and returns a colour beginning with '#'.
func AdjustColour(value string, adjustment int) string {
	i64, err := strconv.ParseInt(value[1:], 16, 32)
	if err != nil {
		return value
	}

	i := int(i64)

	red := (i >> 16) + adjustment
	green := ((i >> 8) & 0x00FF) + adjustment
	blue := (i & 0x0000FF) + adjustment

	red = lo.Max([]int{lo.Min([]int{255, red}), 0})
	green = lo.Max([]int{lo.Min([]int{255, green}), 0})
	blue = lo.Max([]int{lo.Min([]int{255, blue}), 0})

	return fmt.Sprintf("#%.2x%.2x%.2x", red, green, blue)
}

// GetAtIndex returns the value at index i or the default value for type T if it does not exist.
func GetAtIndex[T any](a []T, i int) T {
	if len(a) > i {
		return a[i]
	} else {
		return *new(T)
	}
}
