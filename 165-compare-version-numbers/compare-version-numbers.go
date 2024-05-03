func compareVersion(version1 string, version2 string) int {
	s1, s2 := strings.Split(version1, "."), strings.Split(version2, ".")

	for i := 0; i < min(len(s1), len(s2)); i++ {
		v1, _ := strconv.Atoi(s1[i])
		v2, _ := strconv.Atoi(s2[i])
		if v1 > v2 {
			return 1
		}
		if v2 > v1 {
			return -1
		}
	}

	if len(s1) > len(s2) {
		for i := len(s2); i < len(s1); i++ {
			v, _ := strconv.Atoi(s1[i])
			if v != 0 {
				return 1
			}
		}
	}
	if len(s1) < len(s2) {
		for i := len(s1); i < len(s2); i++ {
			v, _ := strconv.Atoi(s2[i])
			if v != 0 {
				return -1
			}
		}
	}

	return 0
}