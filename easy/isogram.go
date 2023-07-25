package easy

func IsIsogram(word string) bool {
	wordLen := len(word)
	doRepeat := map[string]bool{}
	for i := 0; i < wordLen; i++ {
		if string(word[i]) == " " || string(word[i]) == "-" {
			continue
		}
		// if already exist
		if doRepeat[string(word[i])] {
			return false
		} else {
			doRepeat[string(word[i])] = true
		}
	}
	return true
}
