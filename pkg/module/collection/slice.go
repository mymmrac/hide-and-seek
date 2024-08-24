package collection

import "slices"

func ContainsAll[T comparable](s []T, items ...T) bool {
	for _, item := range items {
		if !slices.Contains(s, item) {
			return false
		}
	}
	return true
}
