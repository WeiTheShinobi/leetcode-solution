func largestMultipleOfThree(digits []int) string {
	count := make([]int, 10)
	sum := 0
	for _, digit := range digits {
		count[digit]++
		sum += digit
	}
	if sum%3 == 1 {
		if count[1] > 0 {
			count[1]--
		} else if count[4] > 0 {
			count[4]--
		} else if count[7] > 0 {
			count[7]--
		} else {
			n := 2
			for i := 2; i <= 8; i += 3 {
				for count[i] > 0 {
					n--
					count[i]--
				}
				if n == 0 {
					break
				}
			}
		}
	}
	if sum%3 == 2 {
		if count[2] > 0 {
			count[2]--
		} else if count[5] > 0 {
			count[5]--
		} else if count[8] > 0 {
			count[8]--
		} else {
			n := 2
			for i := 1; i <= 7; i += 3 {
				for count[i] > 0 {
					n--
					count[i]--
				}
				if n == 0 {
					break
				}
			}
		}
	}

	var result []byte
	for i := 9; i >= 0; i-- {
		for j := 0; j < count[i]; j++ {
			result = append(result, '0'+byte(i))
		}
	}

	if len(result) == 0 {
		return ""
	}
  if result[0] == '0' {
    return "0"
  }
	return string(result)
}