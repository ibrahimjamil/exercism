package easy

func TwoFer(word string) string {
	res := ""

	if len(word) == 0 {
		return "One for you, one for me"
	}

	res = "one for " + word + ", one for me"

	return res
}
