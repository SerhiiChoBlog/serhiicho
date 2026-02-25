package utils

import (
	"fmt"
	"serhii/internal/model"
)

func ExtractIDs[T model.Model](models []T) []int {
	ids := make([]int, len(models))
	for i, m := range models {
		ids[i] = m.Ident()
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
