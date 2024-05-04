func maximumBinaryString(binary string) string {
	count0 := 0
	for i := 0; i < len(binary); i++ {
		if binary[i] == '0' {
			count0++
		}
	}
	if count0 < 2 {
		return binary
	}
	
	firstCount1 := 0
	for i := 0; i < len(binary); i++ {
		if binary[i] == '1' {
			firstCount1++
		} else {
			break
		}
	}

	binary = binary[firstCount1:]
	
	var result []byte
	for i := 0; i < firstCount1; i++ {
		result = append(result, '1')
	}
	for i := 0; i < count0-1; i++ {
		result = append(result, '1')
	}
	result = append(result, '0')
	for i := 0; i < len(binary)-count0; i++ {
		result = append(result, '1')
	}

	return string(result)
}