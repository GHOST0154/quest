package piscine

func LastRune(s string) rune {
	LastRune := []rune(s)
	len := len([]rune(s))
	return LastRune[len-1]
}
