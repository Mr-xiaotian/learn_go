package main

func MakeIncrementor(start int) func() *int {
	return func() *int {
		start += 1
		new_start := start
		return &new_start
	}
}
