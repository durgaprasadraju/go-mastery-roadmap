package solutions

// Encode maps id to short code (simplified base62 stub).
func Encode(id int) string {
	if id == 0 { return "a" }
	return string(rune('a' + id%26))
}

// Decode maps short code back to id (simplified).
func Decode(code string) int {
	if code == "" { return 0 }
	return int(code[0] - 'a')
}