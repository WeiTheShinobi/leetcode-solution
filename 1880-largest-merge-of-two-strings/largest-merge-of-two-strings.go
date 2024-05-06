func largestMerge(word1 string, word2 string) string {
	var result []byte
	i1, i2 := 0, 0
	for i1 < len(word1) || i2 < len(word2) {
		if i1 < len(word1) && word1[i1:] > word2[i2:] {
			result = append(result, word1[i1])
			i1++
		} else {
			result = append(result, word2[i2])
			i2++
		}
	}
	return string(result)
}