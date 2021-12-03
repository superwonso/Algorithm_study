package main

func solution(s string, n int) string {
	str := []rune(s)
	how := rune(n)
	for i, value := range str {
		if 'A' <= value && value <= 'Z' {
			str[i] = ((value + how - 'A') % 26) + 'A'
		} else if 'a' <= value && value <= 'z' {
			str[i] = ((value + how - 'a') % 26) + 'a'
		} else {
			str[i] = value
		}
	}
	return string(str)
}
