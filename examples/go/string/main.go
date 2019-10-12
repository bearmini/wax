package main

func main() {
}

//go:export to_upper
func toUpper(s string) string {
	result := []rune{}
	for _, r := range s {
		//result = append(result, unicode.ToUpper(r))
		result = append(result, r)
	}
	return string(result)
}
