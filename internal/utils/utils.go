package utils

import (
	"fmt"
)

func ExtractIDs[T any](items []T, getID func(T) int) []int {
	ids := make([]int, len(items))
	for i, item := range items {
		ids[i] = getID(item)
	}

	return ids
}

func IntsToStrings(items []int) []string {
	result := make([]string, len(items))

	for i := range items {
		result[i] = fmt.Sprintf("%d", items[i])
	}

	return result
}
