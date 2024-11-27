package piscine


func Capitalize(s string) string {
	variables := []rune(s)
	for i := 0; i < len(variables); i++ {
		if i == 0 || !(variables[i-1] >= 'A' && variables[i-1] <= 'Z') && !(variables[i-1] >= 'a' && variables[i-1] <= 'z') && !(variables[i-1] >= '0' && variables[i-1] <= '9') {
			if variables[i] >= 'a' && variables[i] <= 'z' {
				variables[i] -= 32
			}
		} else {
			if variables[i] >= 'A' && variables[i] <= 'Z' {
				variables[i] += 32
			}
		}
	}
	return string(variables)
}