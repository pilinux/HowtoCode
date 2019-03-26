package iteration

// Repeat - repeat a given string
func Repeat(c string) string {
	var msg string
	for i := 0; i < 5; i++ {
		msg += c
	}

	return msg
}
