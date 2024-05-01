func reversePrefix(word string, ch byte) string {
	var s []byte

	for i := range word {
		s = append(s, word[i])
		if word[i] == ch {
			slices.Reverse(s)
	    return string(s) + word[len(s):]
		}
	}
	
	return word
}