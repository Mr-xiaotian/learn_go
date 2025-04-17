package main

func reverseString(str string) string {
	rune_str := []rune(str)
	reverse_string_rune := []rune{}
	max_str_index := len(rune_str) - 1

	for index := range rune_str {
		reverse_string_rune = append(reverse_string_rune, rune_str[max_str_index-index])
	}
	return string(reverse_string_rune)
}

func isPalindrome(s string) bool {
	return reverseString(s) == s
}

func getLength(ptr_s *string) int {
	return len(*ptr_s)
}
