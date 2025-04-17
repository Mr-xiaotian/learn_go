package main

import "math"

func findMaxKey(m map[string]int) string {
	var max_key string
	max_value := math.MinInt

	for key, value := range m {
		if value <= max_value {
			continue
		}
		max_value = value
		max_key = key
	}
	return max_key
}
